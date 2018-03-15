#! /bin/bash

findpod() {
  local namespace=$1
  local search_query=$2

  local kubectl_args=(--all-namespaces --no-headers)
  local pod_name_field=2
  if [ -n "$namespace" ]; then
    kubectl_args[1]="--namespace=$namespace"
    pod_name_field=1
  fi

  [ -n "$search_query" ] && local fzf_args="--query=$search_query"

  local pod_name=$(kubectl get pod $kubectl_args \
    | awk -v field="$pod_name_field" '{ print $field }' \
    | fzf $fzf_args)
  echo $pod_name
}
