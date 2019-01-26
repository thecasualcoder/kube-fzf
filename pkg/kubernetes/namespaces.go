package kubernetes

import "github.com/arunvelsriram/kube-fzf/pkg/fzf"

// Namespaces represents collection of namespaces
type Namespaces []string

// Filter uses fzf to filter a namespace
func (namespaces Namespaces) Filter(nameQuery string) string {
	result := fzf.Filter(nameQuery, false, namespaces)
	if len(result) != 0 {
		return result[0]
	}
	return ""
}
