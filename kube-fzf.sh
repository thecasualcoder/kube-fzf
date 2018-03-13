#! /bin/bash

findpod() {
  namespace=$1
  search_query=$2

  kubectl_args=(--all-namespaces --no-headers)
  pod_name_field=2
  if [ -n "$namespace" ]; then
    kubectl_args[1]="--namespace=$namespace"
    pod_name_field=1
  fi

  pod_name=$(kubectl get pod $kubectl_args \
    | awk -v field="$pod_name_field" '{ print $field }' \
    | fzf --query $search_query)
  echo $pod_name
}

