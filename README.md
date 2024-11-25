

```markdown
# go-api-test

A simple Go-based API application containerized with Docker and deployable using Kubernetes.

---

## Features

- **Developed in Go**: Efficient and lightweight API application.
- **Containerized with Docker**: Multi-stage build for smaller, secure images.
- **Deployed with Kubernetes**: Easy scalability and management.

---

## Prerequisites

Ensure you have the following installed:
- [Go (1.22 or higher)](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [Kubernetes (kubectl)](https://kubernetes.io/docs/tasks/tools/)
- A local Kubernetes cluster (e.g., [Minikube](https://minikube.sigs.k8s.io/docs/) or [Docker Desktop](https://www.docker.com/products/docker-desktop/)).

---

## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/go-api-test.git
cd go-api-test
```

### 2. Build the Docker Image
Use the provided `Dockerfile` to build the Docker image:
```bash
docker build -t go-api-test:latest .
```

### 3. Run Locally with Docker
Run the application locally:
```bash
docker run -p 8080:8080 go-api-test:latest
```

Access the application at `http://localhost:8080`.

---

## Kubernetes Deployment

### 1. Load the Docker Image into Kubernetes
If you are using Minikube or Docker Desktop, load the image:
```bash
minikube image load go-api-test:latest
```

### 2. Deploy the Application
Apply the Kubernetes manifest:
```bash
kubectl apply -f deployment.yaml
```

Verify the pod is running:
```bash
kubectl get pods
```

### 3. Access the Application
Port-forward to access the API:
```bash
kubectl port-forward deployment/go-api-test 8080:8080
```

Visit `http://localhost:8080`.

Alternatively, expose the deployment with a Kubernetes Service.

---

## Project Structure

```plaintext
📦go-api-test
 ┣ 📂.vscode
 ┃ ┗ 📜launch.json
 ┣ 📂internal
 ┃ ┗ 📂user
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┣ 📜repository.go
 ┃ ┃ ┗ 📜service.go
 ┣ 📂pkg
 ┃ ┗ 📂routes
 ┃ ┃ ┗ 📜router.go
 ┣ 📂tests
 ┃ ┣ 📂integration
 ┃ ┃ ┣ 📜suite_test.go
 ┃ ┃ ┗ 📜user_api_test.go
 ┃ ┗ 📂mocks
 ┃ ┃ ┗ 📜user_repository_mock.go
 ┣ 📜.gitignore
 ┣ 📜Dockerfile
 ┣ 📜LICENSE
 ┣ 📜README.md
 ┣ 📜deployment.yaml
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┣ 📜main.go
 ┗ 📜service.yaml
```

---

## Development

### 1. Install Dependencies
```bash
go mod download
```

### 2. Run Locally
Run the application locally without Docker:
```bash
go run main.go
```

---

## Contributing

Feel free to open issues or submit pull requests for any enhancements or bug fixes.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Contact

For questions or support, please reach out to:
- **Email**: tanvir6@gmail.com
- **GitHub**: [tanvirislam06](https://github.com/tanvirislam06)
```
