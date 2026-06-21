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

Populate `default` namespace with a simple deployment for testing:
```powershell
kubectl create deployment nginx-demo --image=nginx --replicas=3
```

Then verify the pods are running with:
```powershell
kubectl get pods
```

Expected output:
```
NAME                          READY   STATUS    RESTARTS   AGE
nginx-demo-5fd9fc6576-cmxv8   1/1     Running   0          15s       
nginx-demo-5fd9fc6576-mpswb   1/1     Running   0          15s
nginx-demo-5fd9fc6576-rrwvp   1/1     Running   0          15s
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

Expected output for Pods (using `default` as namespace for testing):
```
Listing pods in namespace: default
Found 3 pods
Pod: nginx-demo-5fd9fc6576-cmxv8, Status: Running
Pod: nginx-demo-5fd9fc6576-mpswb, Status: Running
Pod: nginx-demo-5fd9fc6576-rrwvp, Status: Running
```

## Cleanup

```powershell
kubectl delete deployment nginx-demo
kind delete cluster --name kube-pulse-demo
```

Note: running the `kind delete cluster` command may be enough, according to [kubernetes garbage collection docs](https://kubernetes.io/docs/concepts/architecture/garbage-collection/)
- Leaving for now, more research required