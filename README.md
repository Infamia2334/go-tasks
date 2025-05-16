# 🧰 Go Task Management Service

A simple **Task Management Microservice** written in Go, built to demonstrate backend engineering principles and learn Go. The service allows users to **Create** and **Read** tasks, supporting **pagination** and **filtering**.

---

## 📌 Problem Breakdown

The core requirements:

- CRUD(Create and Read for now) operations for tasks.
- Support pagination and filtering in `GET /tasks`.
- Follow microservice architecture principles like single responsibility, loose coupling, and separation of concerns.
- Avoid tight coupling between components to make the system extendable (e.g., for adding users, notifications, etc.).

---

## ⚙️ Design Decisions

### 📁 Project Structure

```bash
go-tasks/
├── cmd/
│ └── server/ # Entry point of the application
│ └── main.go
├── internal/
│ ├── delivery/
│ │ └── http/
│ │   ├── handlers/ # HTTP handlers
│ │   └── routes/ # Route definitions
│ ├── models/ # Task struct and types
│ ├── services/ # Business logic layer
│ └── repository/ # Future data layer with persistant db
├── go.mod
├── Makefile
└── README.md
```
### 🤔 Why This Structure?

- **`cmd/`**: Conventional Go layout for application entry points.
- **`internal/`**: Encapsulates app logic. Follows the Go convention to hide internals.
- **Modular**: Each layer has a single responsibility — easier to test, maintain, and scale.

---

## 🚀 How to Run the Service

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/go-tasks.git
cd go-tasks
```

### 2. Run the service
```bash
go run ./cmd/server
```
Alternatively, use the Makefile:
```bash
make run
```

## 📬 API Documentation

All API endpoints and usage examples are available via Postman.

👉 [View Full API Docs - Postman](https://documenter.getpostman.com/view/13360065/2sB2qWFimy)

---

## 🧩 Microservices Principles Demonstrated

### ✅ Separation of Concerns

- **Handlers** → HTTP-specific logic  
- **Services** → Core business logic  
- **Store** → Simulated persistence (can be swapped with a real DB later)

### ✅ Interface-Based Design

- Services and storage are written as interfaces  
- Enables easy mocking, testing, and replacing components  

### ✅ Scalability-Ready

- Stateless HTTP service — horizontally scalable  
- Easy to containerize using Docker  
- Designed for inter-service communication (REST, gRPC, or message queues)

### ✅ Future-Proof Architecture

- Easily extendable with new microservices like:
  - `UserService`
  - `AuthService`
  - `NotificationService`
- Designed with clear API boundaries  
- Can use REST or gRPC to communicate between services


## 📚 What's Next?
Integrate persistent DB (PostgreSQL, SQLite, etc.)

Add unit tests with testing or testify

Add middleware (e.g., logging, authentication)

Add Docker support and CI/CD

Add metrics and health checks for observability


## 👨‍💻 Author
Dipan — Backend Engineer | Node.js/TS/Python → Go learner 🐹
Feel free to connect or contribute!
