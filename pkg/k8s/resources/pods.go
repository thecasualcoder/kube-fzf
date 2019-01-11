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

// Find finds a pod by name
func (pods Pods) find(name string) *Pod {
	for _, pod := range pods {
		if pod.Name == name {
			return pod
		}
	}

	return nil
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
