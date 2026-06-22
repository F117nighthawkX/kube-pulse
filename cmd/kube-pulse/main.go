package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/F117nighthawkX/kube-pulse/internal/kube"
	"github.com/F117nighthawkX/kube-pulse/internal/kube/health"
	corev1 "k8s.io/api/core/v1"
)

func main() {
	namespace := flag.String("namespace", kube.DefaultNamespace, "Kubernetes namespace to use")
	allNamespaces := flag.Bool("all-namespaces", false, "List pods across all namespaces")

	flag.Parse()

	if *allNamespaces && *namespace != kube.DefaultNamespace {
		log.Fatalf("Cannot use both --namespace and --all-namespaces together")
	}

	ctx := context.Background()

	client, err := kube.CreateNewClient()
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v\n", err)
	}

	var pods []corev1.Pod
	if *allNamespaces {
		pods, err = kube.ListAllPods(ctx, client)
	} else {
		pods, err = kube.ListPods(ctx, client, *namespace)
	}
	if err != nil {
		log.Fatalf("Error listing pods: %v\n", err)
	}

	// for _, pod := range pods {
	// 	fmt.Printf("Pod: %s, Status: %s\n", pod.Name, pod.Status.Phase)
	// }

	healthStatuses := health.AnalyzePods(pods)
	if len(healthStatuses) == 0 {
		fmt.Println("Woops, no pods")
		return
	}

	for _, status := range healthStatuses {
		//fmt.Printf("Pod: %s, Namespace: %s, Ready: %s, Status: %s, Node: %s\n", status.Name, status.Namespace, status.Ready, status.Status, status.Node)
		fmt.Printf("Pod:        %s\n", status.Name)
		fmt.Printf("Namespace:  %s\n", status.Namespace)
		fmt.Printf("Ready:      %s\n", status.Ready)
		fmt.Printf("Status:     %s\n", status.Status)
		fmt.Printf("Restarts:   %d\n", status.Restarts)
		fmt.Printf("Node:       %s\n", status.Node)
		if len(status.Issues) > 0 {
			fmt.Printf("Issues:\n")
			for _, issue := range status.Issues {
				fmt.Printf("  - %s\n", issue)
			}
		}
		fmt.Println("---------------------------------------------------------")
	}
}
