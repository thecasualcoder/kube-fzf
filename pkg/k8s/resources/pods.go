package resources

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Pod represents a Kubernetes Pod
type Pod struct {
	Name string
}

// Pods represents Pod collection
type Pods []Pod

// GetPods gets all pods from the given namespace
func GetPods(clientset *kubernetes.Clientset, namespace string) (Pods, error) {
	podList, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	podItems := podList.Items
	pods := make(Pods, len(podItems))
	for index, pod := range podItems {
		pods[index] = Pod{
			Name: pod.ObjectMeta.Name,
		}
	}

	return pods, nil
}
