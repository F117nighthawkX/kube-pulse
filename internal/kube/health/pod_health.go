package health

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

type PodHealth struct {
	Namespace string
	Name      string
	Ready     string
	Status    string
	Restarts  int32
	Node      string
}

func AnalyzePods(pods []corev1.Pod) []PodHealth {
	var healthStatuses []PodHealth
	for _, pod := range pods {
		health := PodHealth{
			Namespace: pod.Namespace,
			Name:      pod.Name,
			// The default auto-fill from VS Code used `len(pod.Spec.Containers)` here, but that seems to
			// be the desired state defined in the pod spec yaml. Need to use `len(pod.Status.ContainerStatuses)`
			// which gives the runtime status reported by the cluster.
			Ready:    fmt.Sprintf("%d/%d", readyContainers(pod), len(pod.Status.ContainerStatuses)),
			Status:   string(pod.Status.Phase),
			Restarts: countRestarts(pod),
			Node:     pod.Spec.NodeName,
		}
		healthStatuses = append(healthStatuses, health)
		//fmt.Printf("Analyzed pod: %s/%s - Status: %s\n", health.Namespace, health.Name, health.Status)
	}

	return healthStatuses
}

func readyContainers(pod corev1.Pod) int {
	ready := 0
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Ready {
			ready++
		}
	}

	return ready
}

func countRestarts(pod corev1.Pod) int32 {
	var restarts int32
	// ContainerStatus reference: https://pkg.go.dev/k8s.io/api@v0.36.2/core/v1#ContainerStatus
	for _, containerStatus := range pod.Status.ContainerStatuses {
		restarts += containerStatus.RestartCount
	}
	return restarts
}
