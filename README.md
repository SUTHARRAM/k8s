# Go API + React app and deploy it on Kubernetes locally (using Minikube or Docker Desktop’s built-in Kubernetes).

---

### **1. Setup Local Environment**
#### Prerequisites:
- Install [Docker](https://docs.docker.com/get-docker/)
- Install [Kubernetes CLI (kubectl)](https://kubernetes.io/docs/tasks/tools/)
- Choose a local Kubernetes cluster:
  - **Option 1**: [Minikube](https://minikube.sigs.k8s.io/docs/start/) (runs a VM-based cluster)
  - **Option 2**: Docker Desktop (enable Kubernetes in settings)

---

### **2. Project Structure**
Create this folder structure:
```
go-react-k8s/
├── go-api/
│   ├── main.go          # Simple Go API
│   ├── Dockerfile       # Go API container
│   ├── go.mod
├── react-app/
│   ├── src/             # React code
│   ├── Dockerfile       # React container
│   ├── package.json
└── k8s/
    ├── go-deployment.yaml  # Kubernetes configs
    ├── react-deployment.yaml
    ├── service.yaml
```

---

### **3. Build the Go API**
#### Example API (`go-api/main.go`):
```go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go API!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```
#### Go API Dockerfile (`go-api/Dockerfile`):
```dockerfile
# Build stage
FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

---

### **4. Build the React App**
#### Create a simple React app (`react-app/`):
```bash
npx create-react-app react-app
cd react-app
```
Edit `src/App.js` to fetch data from the Go API:
```jsx
import { useEffect, useState } from 'react';

function App() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    fetch('http://go-api-service:8080')  // Kubernetes service name!
      .then(res => res.text())
      .then(data => setMessage(data));
  }, []);

  return <h1>{message || "Loading..."}</h1>;
}

export default App;
```
#### React Dockerfile (`react-app/Dockerfile`):
```dockerfile
# Build stage
FROM node:18 as builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Run stage
FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80
```

---

### **5. Kubernetes Configuration**
#### Kubernetes Deployment/Service Files (`k8s/`):
1. **Go API Deployment (`go-deployment.yaml`)**:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-api
spec:
  replicas: 1
  selector:
    matchLabels:
      app: go-api
  template:
    metadata:
      labels:
        app: go-api
    spec:
      containers:
      - name: go-api
        image: go-api:latest
        ports:
        - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: go-api-service
spec:
  selector:
    app: go-api
  ports:
    - protocol: TCP
      port: 8080
      targetPort: 8080
```

2. **React Deployment (`react-deployment.yaml`)**:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: react-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: react-app
  template:
    metadata:
      labels:
        app: react-app
    spec:
      containers:
      - name: react-app
        image: react-app:latest
        ports:
        - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: react-service
spec:
  type: NodePort  # Expose React to localhost
  selector:
    app: react-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
      nodePort: 30000  # Access via <NodeIP>:30000
```

---

### **6. Build and Deploy**
#### Step 1: Build Docker Images
```bash
# Build Go API image
cd go-api
docker build -t go-api:latest .

# Build React image
cd ../react-app
docker build -t react-app:latest .
```

#### Step 2: Start Kubernetes Cluster
- **Minikube**:
  ```bash
  minikube start
  minikube docker-env  # Use Minikube's Docker daemon
  eval $(minikube -p minikube docker-env)  # Linux/Mac
  ```
- **Docker Desktop**: Ensure Kubernetes is enabled in settings.

#### Step 3: Deploy to Kubernetes
```bash
kubectl apply -f k8s/go-deployment.yaml
kubectl apply -f k8s/react-deployment.yaml
```

#### Step 4: Access the React App
- **Minikube**:
  ```bash
  minikube service react-service --url
  ```
  Open the URL in a browser. You should see "Hello from Go API!" fetched from the backend.

- **Docker Desktop**: Access via `http://localhost:30000`.

---

### **7. Verify**
```bash
kubectl get pods           # Check running pods
kubectl get services       # Verify services
kubectl logs <pod-name>    # Debug issues
```

---

### **Key Concepts Learned**
1. **Docker**: Containerized Go + React apps.
2. **Kubernetes**: Deployments, Services, and inter-pod communication (React → Go API via service name).
3. **Local Development**: Minikube/Docker Desktop for local K8s.

**Next Steps**:
- Add a database (e.g., PostgreSQL with Persistent Volume).
- Use `kubectl port-forward` for debugging.
- Try Helm for package management.


## url - http://192.168.67.2:30000/
