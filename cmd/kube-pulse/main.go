package main

import (
	"fmt"

	"github.com/F117nighthawkX/kube-pulse/internal/kube"
)

func main() {
	fmt.Println("kube-pulse: Kubernetes Resource Health CLI")

	kube.CreateNewClient()
}
