package main

import (
	"github.com/arunvelsriram/kube-fzf/cmd"
	"github.com/arunvelsriram/kube-fzf/cmd/findpod"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var version string

func main() {
	cmd.SetVersion(version)
	findpod.Execute()
}
