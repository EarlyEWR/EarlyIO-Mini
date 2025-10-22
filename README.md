# EarlyIO Mini Learning Environment

A hands-on Go workspace to explore:

- Network programming (TCP, HTTP)
- gRPC service development
- Containerization with Docker & Compose
- Kubernetes deployment basics
- Infrastructure as Code with Terraform
- Orchestration concepts across layers

## Structure
```
examples/network/
  tcp/          # TCP echo server
  tcpclient/    # TCP echo client
  http/         # Basic HTTP server
services/codeconvert/
  server/       # gRPC server (after proto generation)
  client/       # gRPC client (after proto generation)
protos/          # .proto definitions
infra/
  docker/       # Dockerfile + compose
  k8s/          # Kubernetes YAML manifests
  terraform/    # Terraform configs
scripts/         # Helper PowerShell scripts
```

## Prerequisites
- Go 1.22+
- Docker Desktop
- kubectl + a local cluster (kind, k3d, or Docker Desktop Kubernetes)
- Terraform 1.7+
- protoc + plugins:
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

## Generate gRPC Code
```powershell
./scripts/gen-proto.ps1
```
This creates `protos/codeconvert/*.pb.go`. Then uncomment registration & RPC usage lines in gRPC server/client if added later.

## Run Examples
```powershell
# TCP echo server
./scripts/run.ps1 -Target tcp-server
# In another terminal: TCP client
./scripts/run.ps1 -Target tcp-client
# HTTP server
./scripts/run.ps1 -Target http
# gRPC (after proto)
./scripts/run.ps1 -Target grpc-server
./scripts/run.ps1 -Target grpc-client
```

## Docker
Build and run using Compose:
```powershell
cd infra/docker
docker compose up --build
```

## Kubernetes
1. Ensure an image is available (`earlyio/codeconvert:dev`). You can build and load into kind:
```powershell
docker build -f infra/docker/Dockerfile.codeconvert -t earlyio/codeconvert:dev .
kind load docker-image earlyio/codeconvert:dev
```
2. Apply manifests:
```powershell
kubectl apply -f infra/k8s/namespace.yaml
kubectl apply -f infra/k8s/deployment-codeconvert.yaml
kubectl apply -f infra/k8s/service-codeconvert.yaml
# (Optional) ingress
kubectl apply -f infra/k8s/ingress-codeconvert.yaml
```
3. Port-forward to test:
```powershell
kubectl -n earlyio-dev port-forward svc/codeconvert 50051:50051
```

## Terraform (managing simple k8s resources)
```powershell
./scripts/tf-apply.ps1
# Later
./scripts/tf-destroy.ps1
```
Terraform creates the namespace and a configmap as a starting point.

## Orchestration Concepts
- Local processes (scripts) vs Docker Compose multi-service coordination
- Kubernetes for container orchestration & service abstraction
- Terraform for declarative infra state
- gRPC for service-to-service communication; HTTP example for comparison

## Next Steps
- Implement real gRPC methods (uppercase/lowercase) after generation
- Add gateway or REST ↔ gRPC translation
- Add health + readiness probes in Deployment
- Integrate CI (lint, tests) and versioned Docker images
- Expand Terraform to manage secrets or ingress

## Troubleshooting
- Missing protoc: install and ensure it's on PATH
- gRPC compile errors: run `gen-proto.ps1`
- Image not found in cluster: load into kind or push to registry

Happy learning!
