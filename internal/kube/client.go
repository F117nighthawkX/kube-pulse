package kube

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateNewClient() {
	fmt.Println("Creating new Kubernetes client")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error getting user's home directory: %v", err)
	}

	kubeconfigPath := filepath.Join(homeDir, ".kube", "config")
	fmt.Printf("Kubernetes config path: %s\n", kubeconfigPath)
}
