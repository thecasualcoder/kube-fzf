#! /bin/bash

_kube_fzf_usage() {
  cat <<EOF
USAGE:

<function_name> [-n <namespace>] [pod-search-query]

<findpod|tailpod> [-n <namespace>] [pod-search-query]
EOF
}

_kube_fzf_handler() {
  local opt
  while getopts ":hn:" opt; do
    case $opt in
      h)
        _kube_fzf_usage
        return 1
        ;;
      n)
        local namespace="$OPTARG"
        ;;
      \?)
        echo "Invalid Option: -$OPTARG\n"
        _kube_fzf_usage
        return 1
        ;;
      :)
        echo "Option -$OPTARG requires an argument"
        _kube_fzf_usage
        return 1
        ;;
    esac
  done
  shift $((OPTIND - 1))
  [ -n "$1" ] && local pod_search_query=$1

  args="$namespace|$pod_search_query"
}

_kube_fzf_fzf_args() {
  local search_query=$1
  local fzf_args="--height=10 --ansi --reverse"
  [ -n "$search_query" ] && fzf_args="$fzf_args --query=$search_query"
  echo "$fzf_args"
}

_kube_fzf_search_pod() {
  local pod_name
  local namespace=$1
  local pod_search_query=$2
  local pod_fzf_args=$(_kube_fzf_fzf_args "$pod_search_query")

  if [ -z "$namespace" ]; then
    read namespace pod_name <<< \
      $(kubectl get pod --all-namespaces --no-headers \
        | fzf $(printf %s $pod_fzf_args) \
        | awk '{ print $1, $2 }')
  else
    local namespace_fzf_args=$(_kube_fzf_fzf_args "$namespace")
    namespace=$(kubectl get namespaces --no-headers \
      | fzf $(printf %s $namespace_fzf_args) --select-1 \
      | awk '{ print $1 }')

    pod_name=$(kubectl get pod --namespace=$namespace --no-headers \
      | fzf $(printf %s $pod_fzf_args) \
      | awk '{ print $1 }')
  fi

  echo "$namespace|$pod_name"
}

_kube_fzf_echo() {
  local reset_color="\033[0m"
  local bold_green="\033[1;32m"
  local message=$1
  echo -e "\n$bold_green $message $reset_color\n"
}

_kube_fzf_teardown() {
  unset args
  echo $1
}

findpod() {
  local namespace pod_search_query
  _kube_fzf_handler "$@" || return $(_kube_fzf_teardown 1)
  IFS=$'|' read -r namespace pod_search_query <<< "$args"

  local result=$(_kube_fzf_search_pod "$namespace" "$pod_search_query")
  [ -z "$result" ] && return $(_kube_fzf_teardown 1)
  IFS=$'|' read -r namespace pod_name <<< "$result"

  _kube_fzf_echo "kubectl get pod --namespace='$namespace' --output=wide $pod_name"
  kubectl get pod --namespace=$namespace --output=wide $pod_name
  return $(_kube_fzf_teardown 0)
}

tailpod() {
  local namespace pod_search_query
  _kube_fzf_handler "$@" || return $(_kube_fzf_teardown 1)
  IFS=$'|' read -r namespace pod_search_query <<< "$args"

  local result=$(_kube_fzf_search_pod "$namespace" "$pod_search_query")
  IFS=$'|' read -r namespace pod_name <<< "$result"

  local fzf_args=$(_kube_fzf_fzf_args)
  [ -z "$result" ] && return $(_kube_fzf_teardown 1)
  local container_name=$(kubectl get pod $pod_name --namespace=$namespace --output=jsonpath='{.spec.containers[*].name}' \
    | fzf $(printf %s $fzf_args) --select-1)

  _kube_fzf_echo "kubectl logs --namespace='$namespace' --follow $pod_name $container_name"
  kubectl logs --namespace=$namespace --follow $pod_name $container_name
  return $(_kube_fzf_teardown 0)
}

