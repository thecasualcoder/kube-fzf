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
findpod [-a | -n <namespace-query>] [pod-query]
```

### `tailpod`

```
tailpod [-a | -n <namespace-query>] [pod-query]
```

### `execpod`

```
execpod [-a | -n <namespace-query>] [pod-query] <command>
```

### `describepod`

```
describepod [-a | -n <namespace-query>] [pod-query]
```

### Options

```
-a                    -  Search in all namespaces
-n <namespace-query>  -  Find namespaces matching <namespace-query> and do fzf.
                         If there is only one match then it is selected automatically.
-h                    -  Show help
```

## Demo

### `findpod`

![Demo for findpod](/demo/findpod.gif)

### `tailpod`

![Demo for tailpod](/demo/tailpod.gif)

## `tailpod` - multiple containers

![Demo for tailpod with multiple containers](/demo/tailpod-multiple-containers.gif)

### `execpod`

![Demo for execpod](/demo/execpod.gif)

### `execpod` - multiple containers

![Demo for execpod with multiple containers](/demo/execpod-multiple-containers.gif)

### fzf Namespace (only when no match found for the given namespace)

![Demo for wrong namespace](/demo/namespace-matching.gif)

