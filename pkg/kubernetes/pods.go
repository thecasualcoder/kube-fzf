package kubernetes

import (
	"fmt"
	"io"

	"github.com/arunvelsriram/kube-fzf/pkg/fzf"
)

// Pod represents a pod
type Pod struct {
	Name      string
	Namespace string
}

// Pods represents pod collection
type Pods []Pod

// Names returns the names of all the pods
func (pods Pods) Names() []string {
	names := make([]string, len(pods))
	for i, pod := range pods {
		names[i] = pod.Name
	}
	return names
}

// Filter uses fzf to filter one or more pods
func (pods Pods) Filter(nameQuery string, multi bool) Pods {
	filteredPodNames := fzf.Filter(nameQuery, multi, func(in io.WriteCloser) {
		for _, pod := range pods {
			_, _ = fmt.Fprintln(in, pod.Name)
		}
	})
	filteredPods := make(Pods, len(filteredPodNames))
	for i, filteredPodName := range filteredPodNames {
		for _, pod := range pods {
			if pod.Name == filteredPodName {
				filteredPods[i] = pod
				break
			}
		}
	}

	return filteredPods
}

// FilterOne uses fzf to filter a pod
func (pods Pods) FilterOne(nameQuery string) (Pod, error) {
	result := pods.Filter(nameQuery, false)
	if len(result) == 0 {
		return Pod{}, fmt.Errorf("Fzf returned an empty result")
	}

	return result[0], nil
}

// GroupByNamespace returns pods grouped by namespace
func (pods Pods) GroupByNamespace() map[string]Pods {
	grouped := make(map[string]Pods)
	for _, pod := range pods {
		group := grouped[pod.Namespace]
		group = append(group, pod)
		grouped[pod.Namespace] = group
	}

	return grouped
}
