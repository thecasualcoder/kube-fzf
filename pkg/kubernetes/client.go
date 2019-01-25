package kubernetes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Client represents a kubernetes client
type Client struct {
	clientset *kubernetes.Clientset
}

// GetPods gets all pod names from the given namespace
func (c *Client) GetPods(namespace string) ([]string, error) {
	podList, err := c.clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
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

// NewClient creates a kubernetes client
func NewClient(kubeconfig string) (*Client, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		clientset: clientset,
	}, nil
}
