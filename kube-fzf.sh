#!/usr/bin/env bash

_kube_fzf_usage() {
  local func=$1
  echo -e "\nUSAGE:\n"
  case $func in
    findpod)
      cat << EOF
findpod [-a | -n <namespace-query>] [pod-query]

-a                    -  Search in all namespaces
-h                    -  Show help
-n <namespace-query>  -  Find namespaces matching <namespace-query> and do fzf.
                         If there is only one match then it is selected automatically.
EOF
      ;;
    tailpod)
      cat << EOF
tailpod [-a | -n <namespace-query>] [pod-query]

-a                    -  Search in all namespaces
-h                    -  Show help
-n <namespace-query>  -  Find namespaces matching <namespace-query> and do fzf.
                         If there is only one match then it is selected automatically.
EOF
      ;;
    execpod)
      cat << EOF
execpod [-a | -n <namespace-query>] [pod-query] <command>

-a                    -  Search in all namespaces
-h                    -  Show help
-n <namespace-query>  -  Find namespaces matching <namespace-query> and do fzf.
                         If there is only one match then it is selected automatically.
EOF
      ;;
    pfpod)
      cat << EOF
pfpod [ -c | -o | -a | -n <namespace-query>] [pod-query] <source-port:destination-port | port>

-a                    -  Search in all namespaces
-h                    -  Show help
-n <namespace-query>  -  Find namespaces matching <namespace-query> and do fzf.
                         If there is only one match then it is selected automatically.
-o                    -  Open in Browser after port-forwarding
-c                    -  Copy to Clipboard
EOF
      ;;
    describepod)
      cat << EOF
describepod [-a | -n <namespace-query>] [pod-query]

-a                    -  Search in all namespaces
-h                    -  Show help
-n <namespace-query>  -  Find namespaces matching <namespace-query> and do fzf.
                         If there is only one match then it is selected automatically.
EOF
      ;;
  esac
}

_kube_fzf_handler() {
  local opt namespace_query pod_query cmd
  local open=false
  local copy=false
  local OPTIND=1
  local func=$1

  shift $((OPTIND))

  while getopts ":hn:aoc" opt; do
    case $opt in
      h)
        _kube_fzf_usage "$func"
        return 1
        ;;
      n)
        namespace_query="$OPTARG"
        ;;
      a)
        namespace_query="--all-namespaces"
        ;;
      o)
        open=true
        ;;
      c)
        copy=true
        ;;
      \?)
        echo "Invalid Option: -$OPTARG."
        _kube_fzf_usage "$func"
        return 1
        ;;
      :)
        echo "Option -$OPTARG requires an argument."
        _kube_fzf_usage "$func"
        return 1
        ;;
    esac
  done

  shift $((OPTIND - 1))
  if [ "$func" = "execpod" ] || [ "$func" = "pfpod" ]; then
    if [ $# -eq 1 ]; then
      cmd=$1
      [ -z "$cmd" ] && cmd="$SHELL"
    elif [ $# -eq 2 ]; then
      pod_query=$1
      cmd=$2
      if [ -z "$cmd" ]; then
        if [ "$func" = "execpod" ]; then
          echo "Command required." && _kube_fzf_usage "$func" && return 1
        elif [ "$func" = "pfpod" ]; then
          echo "Port required." && _kube_fzf_usage "$func" && return 1
        fi
      fi
    else
      if [ -z "$cmd" ]; then
        if [ "$func" = "execpod" ]; then
          cmd="$SHELL"
        elif [ "$func" = "pfpod" ]; then
          echo "Port required." && _kube_fzf_usage "$func" && return 1
        fi
      fi
    fi
  else
    pod_query=$1
  fi

  args="$namespace_query|$pod_query|$cmd|$open|$copy"
}

_kube_fzf_fzf_args() {
  local search_query=$1
  local extra_args=$2
  local fzf_args="--height=10 --ansi --reverse $extra_args"
  [ -n "$search_query" ] && fzf_args="$fzf_args --query=$search_query"
  echo "$fzf_args"
}

_kube_fzf_search_pod() {
  local namespace pod_name
  local namespace_query=$1
  local pod_query=$2
  local pod_fzf_args=$(_kube_fzf_fzf_args "$pod_query")

  if [ -z "$namespace_query" ]; then
      context=$(kubectl config current-context)
      namespace=$(kubectl config get-contexts --no-headers $context \
        | awk '{ print $5 }')

      namespace=${namespace:=default}
      pod_name=$(kubectl get pod --namespace=$namespace --no-headers \
          | fzf $(echo $pod_fzf_args) \
        | awk '{ print $1 }')
  elif [ "$namespace_query" = "--all-namespaces" ]; then
    read namespace pod_name <<< $(kubectl get pod --all-namespaces --no-headers \
        | fzf $(echo $pod_fzf_args) \
      | awk '{ print $1, $2 }')
  else
    local namespace_fzf_args=$(_kube_fzf_fzf_args "$namespace_query" "--select-1")
    namespace=$(kubectl get namespaces --no-headers \
        | fzf $(echo $namespace_fzf_args) \
      | awk '{ print $1 }')

    namespace=${namespace:=default}
    pod_name=$(kubectl get pod --namespace=$namespace --no-headers \
        | fzf $(echo $pod_fzf_args) \
      | awk '{ print $1 }')
  fi

  [ -z "$pod_name" ] && return 1

  echo "$namespace|$pod_name"
}

_kube_fzf_search_deploy() {
  local namespace deploy_name
  local namespace_query=$1
  local deploy_query=$2
  local deploy_fzf_args=$(_kube_fzf_fzf_args "$deploy_query")

  if [ -z "$namespace_query" ]; then
      context=$(kubectl config current-context)
      namespace=$(kubectl config get-contexts --no-headers $context \
        | awk '{ print $5 }')

      namespace=${namespace:=default}
      deploy_name=$(kubectl get deployment --namespace=$namespace --no-headers \
          | fzf $(echo $deploy_fzf_args) \
        | awk '{ print $1 }')
  elif [ "$namespace_query" = "--all-namespaces" ]; then
    read namespace deploy_name <<< $(kubectl get deployment --all-namespaces --no-headers \
        | fzf $(echo $deploy_fzf_args) \
      | awk '{ print $1, $2 }')
  else
    local namespace_fzf_args=$(_kube_fzf_fzf_args "$namespace_query" "--select-1")
    namespace=$(kubectl get namespaces --no-headers \
        | fzf $(echo $namespace_fzf_args) \
      | awk '{ print $1 }')

    namespace=${namespace:=default}
    deploy_name=$(kubectl get deployment --namespace=$namespace --no-headers \
        | fzf $(echo $deploy_fzf_args) \
      | awk '{ print $1 }')
  fi

  [ -z "$deploy_name" ] && return 1

  echo "$namespace|$deploy_name"
}

_kube_fzf_search_job() {
  local namespace job_name
  local namespace_query=$1
  local job_query=$2
  local job_fzf_args=$(_kube_fzf_fzf_args "$job_query")

  if [ -z "$namespace_query" ]; then
      context=$(kubectl config current-context)
      namespace=$(kubectl config get-contexts --no-headers $context \
        | awk '{ print $5 }')

      namespace=${namespace:=default}
      job_name=$(kubectl get job --namespace=$namespace --no-headers \
          | fzf $(echo $job_fzf_args) \
        | awk '{ print $1 }')
  elif [ "$namespace_query" = "--all-namespaces" ]; then
    read namespace job_name <<< $(kubectl get job --all-namespaces --no-headers \
        | fzf $(echo $job_fzf_args) \
      | awk '{ print $1, $2 }')
  else
    local namespace_fzf_args=$(_kube_fzf_fzf_args "$namespace_query" "--select-1")
    namespace=$(kubectl get namespaces --no-headers \
        | fzf $(echo $namespace_fzf_args) \
      | awk '{ print $1 }')

    namespace=${namespace:=default}
    job_name=$(kubectl get job --namespace=$namespace --no-headers \
        | fzf $(echo $job_fzf_args) \
      | awk '{ print $1 }')
  fi

  [ -z "$job_name" ] && return 1

  echo "$namespace|$job_name"
}

_kube_fzf_echo() {
  local reset_color="\033[0m"
  local bold_green="\033[1;32m"
  local message=$1
  echo -e "\n$bold_green $message $reset_color\n"
}

