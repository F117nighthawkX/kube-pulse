package kube

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const DefaultNamespace = "default"

func ListPods(ctx context.Context, client kubernetes.Interface) ([]corev1.Pod, error) {
	podList, err := client.CoreV1().Pods(DefaultNamespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error listing pods: %v", err)
	}

	fmt.Printf("Listing pods in namespace: %s\n", DefaultNamespace)
	return podList.Items, nil
}
