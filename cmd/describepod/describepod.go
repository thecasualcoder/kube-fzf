package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thecasualcoder/kube-fzf/cmd"
	"github.com/thecasualcoder/kube-fzf/pkg/kubectl"
	"github.com/thecasualcoder/kube-fzf/pkg/kubernetes"
)

var allNamespaces bool
var namespaceName string

var rootCmd = &cobra.Command{
	Use:   "describepod [pod-query]",
	Short: "Describe a pod interactively",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var podName string
		if len(args) == 1 {
			podName = strings.TrimSpace(args[0])
		}

		kubeconfig := viper.GetString("kubeconfig")

		client, err := kubernetes.NewClient(kubeconfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if allNamespaces {
			pods, err := client.GetAllPods()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			filteredPod, err := pods.FilterOne(podName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			kubectl.DescribePod(kubeconfig, filteredPod)
		} else {
			namespaces, err := client.GetNamespaces()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			filteredNamespace, err := namespaces.FilterOne(namespaceName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			pods, err := client.GetPods(filteredNamespace)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			filteredPod, err := pods.FilterOne(podName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			kubectl.DescribePod(kubeconfig, filteredPod)
		}
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
	cobra.OnInitialize(cmd.InitKubeconfig)
	rootCmd.AddCommand(cmd.VersionCmd)
	rootCmd.Flags().BoolVarP(&allNamespaces, "all-namespaces", "a", false, "consider all namespaces")
	rootCmd.Flags().StringVarP(&namespaceName, "namespace", "n", "default", "namespace query")
	rootCmd.Flags().StringP("kubeconfig", "", "", "path to kubeconfig file (default is $HOME/.kube/config)")
	viper.BindPFlag("kubeconfig", rootCmd.Flags().Lookup("kubeconfig"))
	viper.AutomaticEnv()
}
