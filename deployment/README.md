# Kubernetes Deployment (UNTESTED)


## Prerequisites
- [Docker](https://docs.docker.com/get-docker/)
- [Kubernetes CLI](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/)

## Steps
Use Root directory of the project as the working directory to run all this steps.

1. Starting minikube with default configuration
```bash
minikube start
```

2. Build the docker images
```bash
docker build -t edukita-test-backend:latest ./backend
docker build -t edukita-test-frontend:latest ./frontend
```

3. Load the images to minikube (Optional if minikube is not connected directly to docker)
```bash
minikube image load edukita-test-backend:latest
minikube image load edukita-test-frontend:latest
```

4. Apply the deployment and service files
```bash
kubectl create namespace edukita-test
kubectl apply -f deployment/backend-deployment.yaml -n edukita-test
kubectl apply -f deployment/backend-service.yaml -n edukita-test
```

5. Verify the Deployment
```bash
kubectl get pods -n edukita-test
kubectl get svc -n edukita-test
```

## Thigs to Consider for the future
- Use a proper ingress controller for the frontend service
- Separate backend and frontend inte different namespaces
