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
```powershell
kubectl apply -f .\deployments\
```

Then verify the pods are running with:
```powershell
kubectl get pods
```

Expected output:
```
NAME                                        READY   STATUS             RESTARTS   AGE
broken-nginx-deployment-844c4bb7f5-j2qgn    0/1     ImagePullBackOff   0          15m
healthy-nginx-deployment-796d6c889c-gm7mh   1/1     Running            0          15m
healthy-nginx-deployment-796d6c889c-kcmnf   1/1     Running            0          15m
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
Found 3 pods
Pod: broken-nginx-deployment-844c4bb7f5-j2qgn, Namespace: default, Ready: 0/1, Status: Pending, Node: kube-pulse-demo-control-plane
Pod: healthy-nginx-deployment-796d6c889c-gm7mh, Namespace: default, Ready: 1/1, Status: Running, Node: kube-pulse-demo-control-plane
Pod: healthy-nginx-deployment-796d6c889c-kcmnf, Namespace: default, Ready: 1/1, Status: Running, Node: kube-pulse-demo-control-plane
```

## Status Checking

To see detailed information on a pod:
```powershell
kubectl describe pod <pod-name>
kubectl logs <pod-name>
```

Watch the pods update live (seems to work the same as `htop` for Linux resources):
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