package resources

import (
	"fmt"
	"io"

	"github.com/arunvelsriram/kube-fzf/pkg/fzf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Pod represents a Kubernetes Pod
type Pod struct {
	Name string
}

// Pods represents Pod collection
type Pods []*Pod

func (pods Pods) find(name string) *Pod {
	found := pods.findAll([]string{name})
	if len(found) == 0 {
		return nil
	}
	return found[0]
}

func (pods Pods) findAll(names []string) Pods {
	var found Pods
	for _, name := range names {
		for _, pod := range pods {
			if pod.Name == name {
				found = append(found, pod)
				break
			}
		}
	}
	return found
}

// FilterOne filters one pod by name
func (pods Pods) FilterOne(nameQuery string) *Pod {
	filteredPodName := fzf.FilterOne(nameQuery, func(in io.WriteCloser) {
		for _, pod := range pods {
			fmt.Fprintln(in, pod.Name)
		}
	})
	filteredPod := pods.find(filteredPodName)
	return filteredPod
}

// FilterMany filters many pods by name
func (pods Pods) FilterMany(nameQuery string) Pods {
	filteredPodNames := fzf.FilterMany(nameQuery, func(in io.WriteCloser) {
		for _, pod := range pods {
			fmt.Fprintln(in, pod.Name)
		}
	})
	filteredPods := pods.findAll(filteredPodNames)
	return filteredPods
}

// GetPods gets all pods from the given namespace
func GetPods(clientset *kubernetes.Clientset, namespace string) (Pods, error) {
	podList, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	podItems := podList.Items
	pods := make(Pods, len(podItems))
	for index, pod := range podItems {
		pods[index] = &Pod{
			Name: pod.ObjectMeta.Name,
		}
	}

	return pods, nil
}
