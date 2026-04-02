# Service Description

Simple containerized web app packaged with Helm Chart for GitOps friendly Kubernetes deployment.

Web app is built using Go and serves a web page and JSON API showing all the environment variables.
Language choice in a particular case justified by the image size and build time. 
Estimated image size is 10MB, whereas the same app built in Python would be ~100MB + overhead of the runtime.

## Prerequisites

- Docker
- kind
- kubectl
- Helm 3

## Docker Image

Published at: [dreamliner/internal-service](https://hub.docker.com/r/dreamliner/internal-service)

## Setup instructions

### Build the image locally

```bash
docker build -t dreamliner/internal-service:latest .
```

### Create a kind cluster with 3 worker nodes

```bash
kind create cluster --name dev --config kind-config.yaml
```

### Verify nodes are ready

```bash
kubectl get nodes
```

## How to deploy to kind

### Deploy dev

```bash
helm install dev ./helm/internal-service -f ./helm/internal-service/values-dev.yaml
```

### Deploy prod

```bash
helm install prod ./helm/internal-service -f ./helm/internal-service/values-prod.yaml
```

### Verify pods are running

```bash
kubectl get pods -o wide
```

### Test the service

```bash
kubectl port-forward svc/dev-internal-service 8080:8080
curl localhost:8080
```

### Uninstall

```bash
helm uninstall dev
helm uninstall prod
```

## Helm Values Reference

| Key | Description | Default |
|-----|-------------|---------|
| hub | Registry hostname | docker.io |
| image | Image name | dreamliner/internal-service |
| tag | Image tag | latest |
| prod | Production mode flag | false |
| replicaCount | Number of pod replicas | 3 |
| env | Additional environment variables | {} |
| resources.requests.cpu | CPU request | 50m |
| resources.requests.memory | Memory request | 64Mi |
| resources.limits.cpu | CPU limit | 200m |
| resources.limits.memory | Memory limit | 128Mi |
| autoscaling.minReplicas | HPA min replicas | 3 |
| autoscaling.maxReplicas | HPA max replicas | 10 |
| autoscaling.targetCPUUtilizationPercentage | HPA CPU threshold | 70 |
| service.port | Service port | 8080 |

## How to render and apply with Kustomize

### Regenerate base manifests from Helm

```bash
helm template base ./helm/internal-service > gitops/base/all.yaml
```

### Apply dev overlay

```bash
kubectl apply -k gitops/overlays/dev
```

### Apply prod overlay

```bash
kubectl apply -k gitops/overlays/prod
```

## Project Structure

```
.
├── Dockerfile
├── main.go
├── go.mod
├── readme.md
├── kind-config.yaml
├── helm/
│   └── internal-service/
│       ├── Chart.yaml
│       ├── values.yaml
│       ├── values-dev.yaml
│       ├── values-prod.yaml
│       └── templates/
│           ├── _helpers.tpl
│           ├── deployment.yaml
│           ├── service.yaml
│           └── hpa.yaml
└── gitops/
    ├── base/
    │   ├── kustomization.yaml
    │   └── all.yaml
    └── overlays/
        ├── dev/
        │   ├── kustomization.yaml
        │   └── patch-env.yaml
        └── prod/
            ├── kustomization.yaml
            └── patch-env.yaml
```
