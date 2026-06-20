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

Expected output:
```
Connected to Kubernetes server version: v1.36.1
```
Note: the specific version shouldn't matter