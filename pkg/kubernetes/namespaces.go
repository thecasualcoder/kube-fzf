package kubernetes

import (
	"fmt"
	"io"

	"github.com/thecasualcoder/kube-fzf/pkg/fzf"
)

// Namespaces represents collection of namespaces
type Namespaces []string

// FilterOne uses fzf to filter a namespace
func (namespaces Namespaces) FilterOne(nameQuery string) (string, error) {
	result := fzf.Filter(nameQuery, false, func(in io.WriteCloser) {
		for _, namespace := range namespaces {
			_, _ = fmt.Fprintln(in, namespace)
		}
	})
	if len(result) == 0 {
		return "", fmt.Errorf("Fzf returned an empty result")
	}

	return result[0], nil
}
