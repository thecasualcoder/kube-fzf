package kubernetes

import (
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
	filteredPodNames := fzf.Filter(nameQuery, multi, pods.Names())
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
