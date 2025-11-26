# CFX-SRE 

This repository contains two backend applications (Golang and Node.js) and their Kubernetes deployment manifests. It also includes a CI/CD pipeline using GitHub Actions to automate testing, building, and deployment.

---

## Repository Structure

- **backend-go/**: Golang backend application.
- **backend-js/**: Node.js backend application.
- **kustomize/**: Kubernetes manifests managed with Kustomize for simplified deployment.

---

## Running Applications Locally

### Golang Backend
To run the Golang backend locally:

```bash
cd backend-go
go mod download
go run main.go
```

### NodeJS Backend
To run the NodeJS backend locally:

```bash
cd backend-js
npm install
node app.js
```

---

## CI/CD Pipeline

The repository uses GitHub Actions for CI/CD with the following jobs:

### Test
Runs unit tests and verifies the code quality.

### Build Docker Image
Builds Docker images for the backend applications.
The Docker images are pushed to GitHub Container Registry (GHCR) for use in Kubernetes deployments.

### Deploy to Kubernetes
Deploys the applications to the Kubernetes cluster using manifests from the `kustomize/` folder.
Kustomize is used to simplify manifest management.

## Kubernetes Deployment

- The applications are deployed to AWS EKS.

- GitHub Actions authenticates with AWS using an OIDC role `arn:aws-cn:iam::123456789100:role/github-actions-role` for secure access.

- Manifests are managed with Kustomize for environment-specific overlays and simplified deployment.

- `backend-go` will have hostname `backend-go.cfx-sre.com` in production.

- `backend-js` will have hostname `backend-js.cfx-sre.com` in production.
### Notes

- Ensure you have the required AWS permissions for OIDC authentication before running deployments.

- Kustomize simplifies maintaining Kubernetes manifests across multiple environments (e.g., staging, production).

---

### Installing Monitoring Stack

install prometheus stack
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack --version 79.8.2
```

install loki for log database
```
helm repo add grafana https://grafana.github.io/helm-charts

helm install loki grafana/loki --version 6.46.0
```

install promtail for sending application logs to loki
```
helm repo add grafana https://grafana.github.io/helm-charts

helm install promtail grafana/promtail --version 6.17.1
```