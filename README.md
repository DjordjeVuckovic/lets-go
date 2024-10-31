# Let's go

- Let's go is go microservice starter kit. 
- It is a simple and easy to use starter kit for go microservices.

## Folder Structure
### In case of microservice per repository, the folder structure is as follows:
```
/your-microservice-app
│
├── /cmd/
│   └── /service-name/          # Main entry point for your service (e.g., main.go)
│       └── main.go
│
├── /internal/                  # Private application and business logic
│   ├── /config/                # Configuration loading (e.g., config.go)
│   ├── /handlers/              # HTTP handlers, e.g., REST, gRPC, GraphQL
│   ├── /services/              # Business logic, domain services
│   ├── /repository/            # Database access layer
│   └── /middlewares/           # Custom middlewares for HTTP/gRPC
│
├── /pkg/                       # Reusable code for multiple microservices
│   ├── /logger/                # Logging setup (e.g., with Zap or Logrus)
│   └── /utils/                 # Utility functions, helpers
│
├── /api/                       # API specifications (e.g., OpenAPI, Swagger, Protobuf)
│   └── /protobufs/             # Protobuf definitions (if using gRPC)
│
├── /deployments/               # Kubernetes, Docker, Helm, etc. deployment manifests
│   └── /k8s/                   # Kubernetes manifests (e.g., deployment.yaml)
│
├── /manifests/                 # Optional: Kustomize or Kubernetes manifests
│
├── /scripts/                   # Any helper scripts (CI/CD scripts, migration scripts)
│
├── /migrations/                # Database migrations (e.g., SQL files)
│
├── go.mod                      # Module definition
├── go.sum                      # Module dependencies
└── README.md                   # Documentation
```
cmd/: Contains the entry point for your microservice. If you have multiple services or commands, they can each have their own directory inside cmd/. Each service would typically have its own main.go here.

internal/: Holds private code specific to this application. The internal directory prevents other modules from importing its contents. Within this folder:

config/ handles application configuration.
handlers/ is for HTTP or gRPC request handlers.
services/ contains business logic.
repository/ interfaces with databases or other external storage systems.
middlewares/ for custom middleware logic (e.g., authentication, logging).
pkg/: Reusable libraries that might be used in multiple microservices. For example, logging utilities, error handling, and helper functions.

api/: This folder contains API contracts like OpenAPI/Swagger definitions or Protocol Buffers (if using gRPC).

deployments/: Deployment configurations for cloud platforms like Kubernetes (e.g., Helm charts, Kustomize manifests, or Kubernetes YAML files).

scripts/: Helper scripts that automate tasks such as building, testing, running migrations, or deployment.

migrations/: SQL or database migration files (if your service uses a database).

go.mod and go.sum: Go module files to manage dependencies.
### In case of monorepo(multiple microservices in same repository), the folder structure is as follows:
```
/your-monorepo
│
├── /services/                       # Each microservice
│   ├── /service1/                   # First microservice
│   │   ├── /cmd/
│   │   ├── /internal/
│   │   ├── /pkg/
│   │   ├── /api/
│   │   └── go.mod                   # Separate Go module for service1
│   │
│   ├── /service2/                   # Second microservice
│   │   ├── /cmd/
│   │   ├── /internal/
│   │   ├── /pkg/
│   │   ├── /api/
│   │   └── go.mod                   # Separate Go module for service2
│   │
│   └── /service3/                   # Third microservice
│       ├── /cmd/
│       ├── /internal/
│       ├── /pkg/
│       ├── /api/
│       └── go.mod                   # Separate Go module for service3
│
├── /pkg/                            # Shared libraries and code across services
│   ├── /auth/                       # Authentication logic
│   ├── /logger/                     # Logging utilities
│   ├── /middleware/                 # Common HTTP/gRPC middlewares
│   └── go.mod                       # Shared Go module for reusable code
│
├── /deployments/                    # Global deployment configurations (K8s, Docker)
│   └── /k8s/                        # Kubernetes manifests for all services
│
├── /scripts/                        # Scripts for CI/CD or automation
│
├── go.work                           # Go workspace file to manage multiple modules
├── README.md                        # Project-level documentation
└── go.mod                           # Root Go module if needed (optional)
```
