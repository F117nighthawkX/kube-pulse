package kube

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const DefaultNamespace = "default"

func ListPods(ctx context.Context, client kubernetes.Interface, namespace string) ([]corev1.Pod, error) {
	if client == nil {
		return nil, fmt.Errorf("Kubernetes client not initialized")
	}

	if namespace == "" {
		namespace = DefaultNamespace
	}

	podList, err := client.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})

	if err != nil {
		return nil, fmt.Errorf("error listing pods: %v", err)
	}

	return podList.Items, nil
}

func ListAllPods(ctx context.Context, client kubernetes.Interface) ([]corev1.Pod, error) {
	if client == nil {
		return nil, fmt.Errorf("Kubernetes client not initialized")
	}

	// NamespaceAll and NamespaceNone both evaluate to "", but it seems like Kubernetes understands
	// the context of the resource being queried
	podList, err := client.CoreV1().Pods(metav1.NamespaceAll).List(ctx, metav1.ListOptions{})

	if err != nil {
		return nil, fmt.Errorf("error listing pods: %v", err)
	}

	return podList.Items, nil
}
