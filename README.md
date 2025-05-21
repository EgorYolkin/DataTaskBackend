# DataTask: Backend – Your Task Management API

## 🚀 Introduction

**DataTask Backend** is a robust and efficient API server for your task management application. Built with the Go programming language and the Gin framework, this repository contains all the code needed to power the backend of the DataTask application, enabling seamless task, project, and deadline management.

The backend provides a RESTful API that integrates seamlessly with the [**DataTask Frontend**](https://github.com/yourusername/datatask-frontend) (link to your frontend repository), handling data storage, retrieval, and business logic for a comprehensive task management solution.

---

## ✨ Key Backend Features

- **RESTful API:** Provides endpoints for managing users, projects, tasks, kanban boards, and comments, as defined in the Swagger documentation.
- **High Performance:** Leverages Go’s concurrency model and Gin’s lightweight routing for fast response times.
- **Flexible Data Management:** Supports task categorization, prioritization, kanban board organization, and project hierarchies.
- **Secure and Scalable:** Implements API key-based authentication (via Authorization header) and is designed to handle growing user demands.
- **Database Integration:** Compatible with PostgreSQL (as indicated by the example DATABASE_URL in the .env setup).
- **Modern Technology Stack:** Utilizes Go, Gin, and libraries like swaggo/swag for Swagger documentation.

---

## 🖥️ API Endpoints Overview

The API is versioned under `/api/v1` and includes endpoints for:

- **Authentication:** Login (`/auth/login`) and user management (`/user/*`).
- **Projects:** Create, read, update, delete (CRUD) projects (`/project/*`), manage subprojects, and handle user invitations/permissions.
- **Tasks:** CRUD tasks (`/task/*`), assign users to tasks, and retrieve tasks by project, kanban, or user.
- **Kanban Boards:** CRUD kanban boards (`/kanban/*`) and retrieve tasks by kanban.
- **Comments:** Create and retrieve comments for tasks (`/comment/forTask/*`).

*Full API documentation is available in the `swagger.yaml` or `swagger.json` files, or access the interactive Swagger UI at `/swagger/index.html` when the server is running.*

---

## 🛠️ Installation and Setup

To run DataTask Backend locally, follow these steps:

### Running via Docker

```bash
docker compose -f infra/docker/docker-compose.yml --env-file configs/.env.docker up 
```

### Prerequisites

Ensure you have [Go](https://golang.org/) (version 1.18 or later) installed. Optionally, install [Docker](https://www.docker.com/) for containerized deployment. PostgreSQL is recommended for the database.

### Cloning the Repository

```bash
git clone https://github.com/EgorYolkin/DataTaskBackend.git
cd datatask-backend
```

### Installing Dependencies

Use Go modules to install dependencies, including Gin and swaggo/swag for API documentation:

```bash
go mod tidy
```

### Setting Up Environment Variables

Create a `.env` (`.env.docker` for docker build) file in the project’s /infra/config directory and add the following variables:

```text
# Database
DB_USER=user
DB_PASS=pass
DB_HOST=host
DB_PORT=5432
DB_BASE=base

POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=base

# RabbitMQ
RABBITMQ_HOST=host
RABBITMQ_PORT=15672
RABBITMQ_USER=user
RABBITMQ_PASS=pass

# JWT
JWT_SECRET=secret

# Grafana
GF_SECURITY_ADMIN_USER=user
GF_SECURITY_ADMIN_PASSWORD=pass
GF_PORT=3030

# Prometheus
PROMETHEUS_PORT=9090
```

**Important:** Ensure your PostgreSQL database is running and accessible at the specified connection string.

### Running the Application

Run the Go application:

```bash
go run main.go
```

Alternatively, build and run the executable:

```bash
go build -o datatask-backend
./datatask-backend
```

The backend will be available at `http://localhost:8080/api/v1`. Access Swagger documentation at `http://localhost:8080/swagger/index.html` if configured.

## 🧪 Testing

To run the backend tests, use Go’s built-in testing tool:

```bash
go test ./...
```

*Ensure your test database is configured in the `.env` file or use a separate test configuration.*

## 🤝 Contributing

I welcome contributions to the development of DataTask Backend!

If you have suggestions, bug reports, or want to add new features, please:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/AmazingFeature`).
3. Make your changes and commit them (`git commit -m 'Add some AmazingFeature'`).
4. Push to your branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

Please ensure your code adheres to Go coding standards (e.g., `gofmt`) and passes all tests. Update the Swagger documentation (`swagger.yaml`) if you modify or add endpoints.

## 📄 License

This project is licensed under the MIT License. See the LICENSE file for details.