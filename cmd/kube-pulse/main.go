package main

import (
	"context"
	"fmt"

	"github.com/F117nighthawkX/kube-pulse/internal/kube"
)

func main() {
	fmt.Println("\nkube-pulse: Kubernetes Resource Health CLI")

	ctx := context.Background()

	client, err := kube.CreateNewClient()
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %v\n", err)
	}

	pods, err := kube.ListPods(ctx, client, "kube-system")
	if err != nil {
		fmt.Printf("Error listing pods: %v\n", err)
	} else {
		fmt.Printf("Found %d pods\n", len(pods))
	}

	for _, pod := range pods {
		fmt.Printf("Pod: %s, Status: %s\n", pod.Name, pod.Status.Phase)
	}

	version, err := client.Discovery().ServerVersion()
	if err != nil {
		fmt.Printf("Error getting Kubernetes server version: %v\n", err)
	} else {
		fmt.Printf("Connected to Kubernetes server version: %s\n", version.String())
	}
}
