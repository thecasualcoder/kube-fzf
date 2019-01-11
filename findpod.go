package main

import (
	"github.com/arunvelsriram/kube-fzf/cmd/findpod"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	findpod.Execute()
}
