# overview

When building an HTTPS API server project in Go, following a well-structured and modular project layout is critical for maintainability, scalability, and collaboration. 

Go does not enforce a specific project structure, but there are common best practices used by the Go community. 

Below is an example structure and an explanation of each component.


# recommended structure 
https://medium.com/@gaurang.m/golang-project-structure-13b953fb6117

my-go-api-server/
├── cmd/
│   └── api/
│       └── main.go          # Entry point for the application
├── config/
│   └── config.go            # Configuration handling
├── internal/
│   ├── handlers/            # HTTP handlers for routes
│   │   ├── user.go          # Example: User-related handlers
│   │   └── health.go        # Example: Health check handler
│   ├── services/            # Business logic
│   │   └── user_service.go  # Example: User-related business logic
│   ├── repositories/        # Data access layer
│   │   └── user_repo.go     # Example: Database queries for users
│   ├── middleware/          # HTTP middleware
│   │   └── auth.go          # Example: Authentication middleware
│   └── models/              # Data models
│       └── user.go          # Example: User data structure
├── pkg/
│   ├── logger/              # Utilities for logging
│   │   └── logger.go
│   ├── utils/               # General helper functions
│       └── utils.go
├── docs/
│   └── api-spec.yaml        # API specifications (e.g., OpenAPI/Swagger)
├── test/
│   ├── integration/         # Integration tests
│   │   └── user_integration_test.go
│   └── unit/                # Unit tests
│       └── user_service_test.go
├── go.mod                   # Module definition file
└── README.md                # Project documentation

# directories

1. cmd/:

Contains the entry points for your application.
The subdirectory (e.g., api) holds the main.go file for your server. If you build multiple binaries, each binary gets its own subdirectory.

2. config/:

Contains code for loading and managing application configuration (e.g., from environment variables or configuration files).

3. internal/:

This is where your application’s core logic resides. It is private to the module (not accessible to external modules).
Subdirectories:

handlers/: HTTP handlers (controller layer) to process API requests.
services/: Business logic and rules (service layer).
repositories/: Data access layer for interacting with databases or other storage systems.
middleware/: Middleware for handling cross-cutting concerns like logging, authentication, or rate limiting.
models/: Data structures representing your domain models.

4. pkg/:

Holds reusable code that can be imported by other parts of the project or other projects (e.g., logger, utility functions).
This directory is optional if your project doesn’t need reusable packages.

5. docs/:

API documentation, such as OpenAPI/Swagger specs.

6. test/:

Dedicated space for test files.
You can separate unit tests and integration tests for better organization.

7. go.mod:

Defines your module and manages dependencies.

8. README.md:

Provides documentation about your project, how to run it, and its purpose.