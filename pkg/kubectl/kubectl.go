package kubectl

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/arunvelsriram/kube-fzf/pkg/kubernetes"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubernetes/pkg/kubectl/cmd/describe"
	"k8s.io/kubernetes/pkg/kubectl/cmd/get"
	kubectlutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

const (
	getCmd      = "get"
	describeCmd = "describe"
)

type flags map[string]string

func newCmd(name, kubeconfig, namespace string, flags flags) *cobra.Command {
	configFlags := genericclioptions.NewConfigFlags(true)
	configFlags.KubeConfig = &kubeconfig
	configFlags.Namespace = &namespace
	factory := kubectlutil.NewFactory(configFlags)
	ioStreams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	var cmd *cobra.Command
	switch name {
	case getCmd:
		cmd = get.NewCmdGet("kubectl", factory, ioStreams)
	case describeCmd:
		cmd = describe.NewCmdDescribe("kubectl", factory, ioStreams)
	}

	for flag, value := range flags {
		cmd.Flag(flag).Value.Set(value)
	}

	return cmd
}

// GetPods get pods from the kbernetes cluster
func GetPods(kubeconfig string, pods kubernetes.Pods) {
	groupedPods := pods.GroupByNamespace()
	for namespace, pods := range groupedPods {
		cmd := newCmd(getCmd, kubeconfig, namespace, flags{"output": "wide"})
		args := []string{"pod"}
		args = append(args, pods.Names()...)
		fmt.Printf("\nNamespace: %s\n\n", namespace)
		cmd.Run(cmd, args)
	}
}

// DescribePod describes a pod in the kubernetes cluster
func DescribePod(kubeconfig string, pod kubernetes.Pod) {
	cmd := newCmd(describeCmd, kubeconfig, pod.Namespace, flags{"": ""})
	args := []string{"pod", pod.Name}
	cmd.Run(cmd, args)
}
