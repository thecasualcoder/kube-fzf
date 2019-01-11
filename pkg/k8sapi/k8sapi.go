package k8sapi

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetPods gets all pod names from the given namespace
func GetPods(clientset *kubernetes.Clientset, namespace string) ([]string, error) {
	podList, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	podItems := podList.Items
	pods := make([]string, len(podItems))
	for index, pod := range podItems {
		pods[index] = pod.ObjectMeta.Name
	}

	return pods, nil
}
