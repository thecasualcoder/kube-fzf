package findpod

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/arunvelsriram/kube-fzf/pkg/k8s/resources"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var allNamespaces bool
var namespaceName string

var rootCmd = &cobra.Command{
	Use:   "findpod",
	Short: "Finds a pod",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var podName string
		if len(args) == 1 {
			podName = args[0]
		}
		fmt.Println(podName)

		var kubeconfig string
		home, err := homedir.Dir()
		if err != nil {
			panic(err.Error())
		}

		kubeconfig = filepath.Join(home, ".kube", "config")

		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}

		pods, err := resources.GetPods(clientset, "default")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(pods)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&allNamespaces, "all-namespaces", "a", false, "consider all namespaces")
	rootCmd.Flags().StringVarP(&namespaceName, "namespace", "n", "", "namespace pattern")
}
