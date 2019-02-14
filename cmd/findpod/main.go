package main

import (
	"github.com/thecasualcoder/kube-fzf/cmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var version string

func main() {
	cmd.SetVersion(version)
	Execute()
}
