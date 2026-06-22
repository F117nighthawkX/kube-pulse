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
	Issues    []string
}

func AnalyzePods(pods []corev1.Pod) []PodHealth {
	var healthStatuses []PodHealth
	for _, pod := range pods {
		health := PodHealth{
			Namespace: pod.Namespace,
			Name:      pod.Name,
			// The auto-fill from VS Code used `len(pod.Spec.Containers)` here, but that seems to be the
			// desired state defined in the pod spec yaml. Need to use `len(pod.Status.ContainerStatuses)`
			// which gives the runtime status reported by the cluster.
			Ready:    fmt.Sprintf("%d/%d", readyContainers(pod), len(pod.Status.ContainerStatuses)),
			Status:   string(pod.Status.Phase),
			Restarts: countRestarts(pod),
			Node:     pod.Spec.NodeName,
			Issues:   []string{},
		}

		// VS Code auto-fill suggested the `...` elipsis syntax here, which seems to share some similarity
		// to JavaScript's spread operator. Go calls it a variadic function, and it has some notable
		// differences. See https://go101.org/article/function.html
		health.Issues = append(health.Issues, phaseIssues(pod)...)
		health.Issues = append(health.Issues, containerIssues(pod)...)

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

func phaseIssues(pod corev1.Pod) []string {
	var issues []string
	phase := pod.Status.Phase

	if phase == corev1.PodRunning || phase == corev1.PodSucceeded {
		return issues
	}

	// Looks like PodUnknown has been deprecated, will never be set

	issues = append(issues, fmt.Sprintf("Pod is in %s phase", phase))

	return issues
}

func containerIssues(pod corev1.Pod) []string {
	var issues []string
	for _, containerStatus := range pod.Status.ContainerStatuses {
		containerName := containerStatus.Name

		if !containerStatus.Ready {
			msg := fmt.Sprintf("Container %s is not ready", containerName)
			issues = append(issues, msg)
		}

		if containerStatus.RestartCount > 0 {
			msg := fmt.Sprintf("Container %s has restarted %d times", containerName, containerStatus.RestartCount)
			issues = append(issues, msg)
		}

		if containerStatus.State.Waiting != nil {
			msg := fmt.Sprintf("Container %s is waiting: %s", containerName, containerStatus.State.Waiting.Reason)
			issues = append(issues, msg)
		}

		if containerStatus.State.Terminated != nil {
			msg := fmt.Sprintf("Container %s is terminated: %s", containerName, containerStatus.State.Terminated.Reason)
			issues = append(issues, msg)
		}
	}

	return issues
}
