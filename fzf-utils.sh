load_kube_fzf(){
  if [ ${PWD##*/} == "kube-fzf" ]
  then
    source ./kube-fzf.sh
  else
    source kube-fzf.sh
  fi
}
load_kube_fzf