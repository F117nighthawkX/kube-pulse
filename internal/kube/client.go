package kube

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateNewClient(kubeconfigPath string) (kubernetes.Interface, error) {
	//fmt.Println("Creating new Kubernetes client")

	if kubeconfigPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("error getting user's home directory: %v", err)
		}
		kubeconfigPath = filepath.Join(homeDir, ".kube", "config")
	}
	//fmt.Printf("Kubernetes config path: %s\n", kubeconfigPath)

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("error building Kubernetes config: %v", err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error creating Kubernetes client: %v", err)
	}

	//fmt.Printf("Kubernetes client created successfully: %T\n", client)
	return client, nil
}
