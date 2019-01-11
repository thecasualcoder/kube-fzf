package findpod

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/arunvelsriram/kube-fzf/pkg/k8s/resources"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		fmt.Println()
		var podName string
		if len(args) == 1 {
			podName = args[0]
		}
		fmt.Println(podName)

		kubeconfig := viper.GetString("kubeconfig")
		fmt.Println(kubeconfig)
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

func initKubeconfig() {
	if !viper.IsSet("kubeconfig") || viper.GetString("kubeconfig") == "" {
		home, err := homedir.Dir()
		if err != nil {
			panic(err.Error())
		}

		viper.SetDefault("kubeconfig", filepath.Join(home, ".kube", "config"))
	}
}

func init() {
	cobra.OnInitialize(initKubeconfig)
	rootCmd.Flags().BoolVarP(&allNamespaces, "all-namespaces", "a", false, "consider all namespaces")
	rootCmd.Flags().StringVarP(&namespaceName, "namespace", "n", "", "namespace pattern")
	rootCmd.Flags().BoolVarP(&multiSelect, "multi", "m", true, `find multiple pods
use tab/shift+tab to select/de-select from the interactive list`)
	rootCmd.Flags().StringP("kubeconfig", "", "", "path to kubeconfig file (default is $HOME/.kube/config)")
	viper.BindPFlag("kubeconfig", rootCmd.Flags().Lookup("kubeconfig"))
	viper.AutomaticEnv()
}
