# kube-fzf

Shell functions using [`kubectl`](https://kubernetes.io/docs/reference/kubectl/overview/) and [`fzf`](https://github.com/junegunn/fzf) to enable command-line fuzzy searching of [Kubernetes](https://kubernetes.io/) [Pods](https://kubernetes.io/docs/concepts/workloads/pods/pod/). It helps to interactively

- search for a Pod
- tail a container of a Pod
- exec in to a container of a Pod

## Prerequisite

[`fzf`](https://github.com/junegunn/fzf) must be available in the PATH

## Install

```
git clone https://github.com/arunvelsriram/kube-fzf.git ~/.kube-fzf
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

### `findpod`

```
findpod [-n <namespace-query>] [pod-query]

findpod -h for help
```

### `tailpod`

```
tailpod [-n <namespace-query>] [pod-query]

tailpod -h for help
```

### `execpod`

```
execpod [-n <namespace-query>] [pod-query] <command>

execpod -h for help
```

**Note:** If there is only one match for `<namespace-query>` then it is selected automatically.

## Demo

### `findpod`

![Demo for findpod](/demo/findpod.gif)

### `tailpod`

![Demo for tailpod](/demo/tailpod.gif)

### `execpod`
![Demo for execpod](/demo/execpod.gif)

### fzf Namespace (only when no match found for the given namespace)

![Demo for wrong namespace](/demo/namespace.gif)

### fzf Containers inside a Pod

![Demo for fzf containers inside a pod](/demo/containers.gif)

