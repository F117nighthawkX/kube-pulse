# kube-pulse

Go/Kubernetes learning project

## Local Setup

Using kind for Kubernetes testing, with DockerDesktop initiated
```powershell
kind create cluster --name kube-pulse-demo
```

Check the status of the node with:
```powershell
kubectl get nodes
```

Expected output:
```
NAME                            STATUS   ROLES           AGE   VERSION
kube-pulse-demo-control-plane   Ready    control-plane   12m   v1.36.1
```

You can also check the cluster with:
```powershell
kind get clusters
```

Expected output:
```
kube-pulse-demo
```

Populate `default` namespace with deployments from `.\deployments\` (run from project root):
Note: this is nifty, rerunning the command lets you create new deployments and configure existing ones without deleting them first
```powershell
kubectl apply -f .\deployments\
```

Then verify the pods are running with:
```powershell
kubectl get pods
```

Expected output:
```
NAME                                            READY   STATUS             RESTARTS      AGE
broken-command-deployment-f9f6774f-xd6zk        0/1     Completed          2 (19s ago)   20s
broken-image-deployment-5ff68fdb98-xfbjb        0/1     ImagePullBackOff   0             20s
healthy-nginx-deployment-5d88895494-8jclx       1/1     Running            0             20s
healthy-nginx-deployment-5d88895494-n4ms7       1/1     Running            0             20s
missing-resources-deployment-7dc54d449c-wrck4   1/1     Running            0             20s
```

## Run the Project

If you haven't in a while, make sure node, cluster, and pods exist:
```powershell
kubectl get nodes
kind get clusters
kubectl get pods
```

If not, repeat `Local Setup`

Run program from project root:
```powershell
go run .\cmd\kube-pulse
```

Expected output for Pod health statuses (using `default` as namespace for testing):
```
Listing pods in namespace: default
Found 5 pods
Pod:        broken-command-deployment-f9f6774f-v44b9
Namespace:  default
Ready:      0/1
Status:     Running
Node:       kube-pulse-demo-control-plane
-----------------------------
Pod:        healthy-nginx-deployment-5d88895494-fj6hx
Namespace:  default
Ready:      1/1
Status:     Running
Node:       kube-pulse-demo-control-plane
...
```

## Status Checking

To see detailed information on a pod:
```powershell
kubectl describe pod <pod-name>
kubectl logs <pod-name>
```

Watch the pods update live (not the same as `htop` for Linux resources, rather sends updates to new line in terminal):
```powershell
kubectl get pods --watch
```

## Cleanup

Delete demo deployments:
```powershell
kubectl delete -f .\deployments\
```

Delete kind cluster:
```powershell
kind delete cluster --name kube-pulse-demo
```

Note: running the `kind delete cluster` command may be enough, according to [kubernetes garbage collection docs](https://kubernetes.io/docs/concepts/architecture/garbage-collection/)
- Leaving for now, more research required