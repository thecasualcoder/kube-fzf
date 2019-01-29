package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// InitKubeconfig resolves the value of kubeconfig
func InitKubeconfig() {
	if !viper.IsSet("kubeconfig") || viper.GetString("kubeconfig") == "" {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.SetDefault("kubeconfig", filepath.Join(home, ".kube", "config"))
	}
}
