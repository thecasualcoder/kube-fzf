package kubectl

import (
	"os"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubernetes/pkg/kubectl/cmd/get"
	kubectlutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

// GetPods get pods from the kbernetes cluster
func GetPods(kubeconfig, namespace string, podNames []string) {
	configFlags := genericclioptions.NewConfigFlags(true)
	configFlags.KubeConfig = &kubeconfig
	configFlags.Namespace = &namespace
	factory := kubectlutil.NewFactory(configFlags)
	ioStreams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	cmd := get.NewCmdGet("kubectl", factory, ioStreams)
	cmd.Flag("output").Value.Set("wide")
	args := []string{"pod"}
	args = append(args, podNames...)
	cmd.Run(cmd, args)
}
