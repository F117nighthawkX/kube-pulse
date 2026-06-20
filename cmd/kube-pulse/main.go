package main

import (
	"fmt"

	"github.com/F117nighthawkX/kube-pulse/internal/kube"
)

func main() {
	fmt.Println("\nkube-pulse: Kubernetes Resource Health CLI")

	client, err := kube.CreateNewClient()
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %v\n", err)
	}

	version, err := client.Discovery().ServerVersion()
	if err != nil {
		fmt.Printf("Error getting Kubernetes server version: %v\n", err)
	} else {
		fmt.Printf("Connected to Kubernetes server version: %s\n", version.String())
	}
}
