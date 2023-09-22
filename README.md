# kube-fzf

Shell commands using [`kubectl`](https://kubernetes.io/docs/reference/kubectl/overview/) and [`fzf`](https://github.com/junegunn/fzf) for command-line fuzzy searching of [Kubernetes](https://kubernetes.io/) [Pods](https://kubernetes.io/docs/concepts/workloads/pods/pod/). It helps to interactively:

* search for a Pod
* tail a container of a Pod
* exec in to a container of a Pod
* describe a pod
* port forward pod

## Prerequisite

* [`fzf`](https://github.com/junegunn/fzf)
* [`xclip`](https://linux.die.net/man/1/xclip) Only for Linux and it is optional

## Install

### Using Homebrew

```
brew tap thecasualcoder/stable
brew install kube-fzf
```

### Manual

```
git clone https://github.com/bmouser/kube-fzf.git ~/.kube-fzf
./install.sh
```

## Usage

### `findpod`

```
findpod [-a | -n <namespace-query>] [pod-query]
```

### `findeploy`

```
findeploy [-a | -n <namespace-query>] [deploy-query]
```

### `editdeploy`

```
editdeploy [-a | -n <namespace-query>] [deploy-query]
```

### `deletepod`

```
deletepod [-a | -n <namespace-query>] [pod-query]
```

### `describepod`

```
describepod [-a | -n <namespace-query>] [pod-query]
```

### `tailpod`

```
tailpod [-a | -n <namespace-query>] [pod-query]
```

### `taildeploy`

```
taildeploy [-a | -n <namespace-query>] [deployment-query]
```

### `execpod`

```
execpod [-a | -n <namespace-query>] [pod-query] <command>
```

### `pfpod`

```
pfpod [-c | -o | -a | -n <namespace-query>] [pod-query] <port>
```

#### Options

```
-a                    -  Search in all namespaces
-n <namespace-query>  -  Find namespaces matching <namespace-query> and do fzf.
                         If there is only one match then it is selected automatically.
-h                    -  Show help
```

## Demo

### `findpod`

![Demo for findpod](/demo/findpod.gif)

### `describepod`

![Demo for describepod](/demo/describepod.gif)

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

## Similar Projects

For switching Kubernetes contexts and namespaces interactively from the command-line use [`kubectx`](https://github.com/ahmetb/kubectx)
