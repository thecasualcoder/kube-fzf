package kubernetes

import (
	"fmt"

	"github.com/arunvelsriram/kube-fzf/pkg/fzf"
)

// Namespaces represents collection of namespaces
type Namespaces []string

// FilterOne uses fzf to filter a namespace
func (namespaces Namespaces) FilterOne(nameQuery string) (string, error) {
	result := fzf.Filter(nameQuery, false, namespaces)
	if len(result) == 0 {
		return "", fmt.Errorf("Fzf returned an empty result")
	}

	return result[0], nil
}
