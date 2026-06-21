# kube-pulse

Go/Kubernetes learning project

## Local Setup

Using kind for Kubernetes testing

```powershell
kind create cluster --name kube-pulse-demo
kubectl get nodes
```

## Verify Client Connection

Make sure cluster exists:

```powershell
kind get clusters
kubectl get nodes
```

If not, repeat `Local Setup`

Run program from project root:

```powershell
go run .\cmd\kube-pulse
```

Expected output for Pods (using `kube-system` as namespace for testing):
```
Listing pods in namespace: kube-system
Found 8 pods
Pod: coredns-xxx, Status: Running
Pod: coredns-xxx, Status: Running
Pod: etcd-kube-pulse-demo-control-plane, Status: Running
Pod: kindnet-xxx, Status: Running
Pod: kube-apiserver-kube-pulse-demo-control-plane, Status: Running
Pod: kube-controller-manager-kube-pulse-demo-control-plane, Status: Running
Pod: kube-proxy-xxx, Status: Running
Pod: kube-scheduler-kube-pulse-demo-control-plane, Status: Running
```
Note: IDs may differ slightly