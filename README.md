# kube-fzf (WIP)

Shell functions using [`kubectl`](https://kubernetes.io/docs/reference/kubectl/overview/) and [`fzf`](https://github.com/junegunn/fzf) to enable command-line fuzzy searching of [Kubernetes](https://kubernetes.io/) [Pods](https://kubernetes.io/docs/concepts/workloads/pods/pod/). It helps to interactively

- search for a Pod
- tail a container of a Pod
- exec in to a container of a Pod

## Prerequisite

[`fzf`](https://github.com/junegunn/fzf) must be available in the PATH

## Install

```
git clone https://github.com/arunvelsriram/kube-fzf.git --depth=1 ~/.kube-fzf
```
```
# zsh users
echo "[ -f ~/.kube-fzf/kube-fzf.sh ] && source ~/.kube-fzf/kube-fzf.sh" >> ~/.zshrc
source ~/.zshrc
```

```
# bash users
echo "[ -f ~/.kube-fzf/kube-fzf.sh ] && source ~/.kube-fzf/kube-fzf.sh" >> ~/.bashrc
source ~/.bashrc
```

## Usage

```
<function_name> [-n <namespace>] [-p <pod-search-query>] [-c <container-search-query>]
```

```
<getpod|findpod|tailpod> [-n <namespace>] [-p <pod-search-query>] [-c <container-search-query>]
```
