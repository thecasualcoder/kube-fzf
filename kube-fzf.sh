#! /bin/bash

_kube_fzf_handler() {
  local namespace pod_search_query

  while [ $# -gt 0 ]
  do
    local key="$1"
    case $key in
      -n)
        namespace="$2"
        shift
        shift
        ;;
      -p)
        pod_search_query="$2"
        shift
        shift
        ;;
      -c)
        container_search_query="$2"
        shift
        shift
        ;;
      *)
        echo "Unknown argument: $1"
        echo "USAGE:"
        echo "WIP"
        return 1
    esac
  done

  args="$namespace|$pod_search_query|$container_search_query"
  return 0
}

_kube_fzf_fzf_args() {
  local search_query=$1
  local fzf_args="--select-1"
  [ -n "$search_query" ] && fzf_args="$fzf_args --query=$pod_search_query"
  echo "$fzf_args"
}

_kube_fzf_findpod() {
  local namespace=$1
  local pod_search_query=$2
  local namespace_arg="--all-namespaces"
  local pod_name_field=2
  if [ -n "$namespace" ]; then
    namespace_arg="--namespace=$namespace"
    pod_name_field=1
  fi
  local fzf_args=$(_kube_fzf_fzf_args "$pod_search_query")
  local pod_name=$(kubectl get pod $namespace_arg --no-headers \
    | awk -v field="$pod_name_field" '{ print $field }' \
    | fzf $(printf %s $fzf_args))
  echo $pod_name
}

_kube_fzf_echo() {
  local reset_color="\033[0m"
  local bold_green="\033[1;32m"
  local message=$1
  echo -e "\n$bold_green $message $reset_color\n"
}

findpod() {
  local namespace pod_search_query _container_search_query
  _kube_fzf_handler "$@" || return 1
  IFS=$'|' read -r namespace pod_search_query _container_search_query <<< "$args"
  _kube_fzf_findpod "$namespace" "$pod_search_query"
  _kube_fzf_cleanup
  return 0
}

getpod() {
  local namespace pod_search_query _container_search_query
  _kube_fzf_handler "$@" || return 1
  IFS=$'|' read -r namespace pod_search_query _container_search_query <<< "$args"
  local pod_name=$(_kube_fzf_findpod "$namespace" "$pod_search_query")
  _kube_fzf_echo "kubectl get pod --namespace='$namespace' --output wide $pod_name"
  kubectl get pod --namespace=$namespace --output wide $pod_name
  _kube_fzf_cleanup
}

tailpod() {
  local namespace pod_search_query container_search_query
  _kube_fzf_handler "$@" || return 1
  IFS=$'|' read -r namespace pod_search_query container_search_query <<< "$args"
  local pod_name=$(_kube_fzf_findpod "$namespace" "$pod_search_query")
  local fzf_args=$(_kube_fzf_fzf_args "$container_search_query")
  local container_name=$(kubectl get pod $pod_name --output jsonpath='{.spec.containers[*].name}' \
    | fzf $(printf %s $fzf_args))
  _kube_fzf_echo "kubectl logs --namespace=$namespace --follow $pod_name $container_name"
  kubectl logs --namespace=$namespace --follow $pod_name $container_name
  _kube_fzf_cleanup
}

_kube_fzf_cleanup() {
  unset args
}
