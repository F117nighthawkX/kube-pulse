package main

import (
	"context"
	"fmt"
	"log"

	"github.com/F117nighthawkX/kube-pulse/internal/kube"
	"github.com/F117nighthawkX/kube-pulse/internal/kube/health"
)

func main() {
	fmt.Println("\nkube-pulse: Kubernetes Resource Health CLI")

	ctx := context.Background()

	client, err := kube.CreateNewClient()
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v\n", err)
	}

	pods, err := kube.ListPods(ctx, client, "default")
	if err != nil {
		log.Fatalf("Error listing pods: %v\n", err)
	} else {
		fmt.Printf("Found %d pods\n", len(pods))
	}

	// for _, pod := range pods {
	// 	fmt.Printf("Pod: %s, Status: %s\n", pod.Name, pod.Status.Phase)
	// }

	healthStatuses := health.AnalyzePods(pods)
	for _, status := range healthStatuses {
		fmt.Printf("Pod: %s, Namespace: %s, Ready: %s, Status: %s, Node: %s\n", status.Name, status.Namespace, status.Ready, status.Status, status.Node)
	}
}
