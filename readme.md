# 🚀 Go Monolithic Backend API

A clean and scalable monolithic backend API built with **Go**, **PostgreSQL**, and **Docker**.

---

## 📂 Folder Structure

```plaintext
go-monolith-api/
│── cmd/
│   ├── main.go               # Application entry point
│
├── config/
│   ├── config.go             # Configuration setup
│
├── internal/
│   ├── app/
│   │   ├── handler/
│   │   │   ├── user_handler.go
│   │   ├── repository/
│   │   │   ├── user_repository.go
│   │   ├── usecase/
│   │   │   ├── user_usecase.go
│   │   ├── dto/
│   │   │   ├── user_dto.go
│   │   ├── model/
│   │   │   ├── user.go
│   │   ├── middleware/
│   │   │   ├── auth_middleware.go
│   │
│   ├── db/
│   │   ├── db.go             # Database connection
│   │   ├── migrations/       # Database migrations
│   │
│   ├── routes/
│   │   ├── routes.go         # Centralized route registration
│   │   ├── user_routes.go    # User-specific routes
│   │
│   ├── util/
│   │   ├── jwt.go            # JWT Token utils
│   │   ├── hash.go           # Password hashing utils
│
├── test/
│   ├── user_test.go
│
├── .env                      # Environment variables
├── Dockerfile                # Docker configuration
├── docker-compose.yml        # Docker Compose setup
├── go.mod                    # Go module dependencies
├── go.sum                    # Go package checksum
├── Makefile                  # Makefile for easy commands
├── README.md                 # Project documentation
```

---

## 🛠 Prerequisites

Make sure you have installed:
- [Go](https://go.dev/doc/install)
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)

---

## 🏗 Setup & Run

### 🔹 1️⃣ Clone the Repository
```sh
git clone https://github.com/your-repo/go-monolith-api.git
cd go-monolith-api
```

### 🔹 2️⃣ Create a `.env` File
Create a `.env` file and configure the database URL:

```sh
DATABASE_URL="host=postgres user=ashish password=Ashish@#21 dbname=mydatabase port=5432 sslmode=disable"
```

### 🔹 3️⃣ Run with Docker
```sh
docker-compose up -d  # Start PostgreSQL
make run              # Run the application
```

### 🔹 4️⃣ Run Without Docker (Local Setup)
If you want to run it without Docker:

```sh
export $(cat .env | xargs)  # Load env variables
make migrate                # Run database migrations
make run                    # Start the Go application
```

---

## 🎯 Makefile Commands

📂 **Makefile** provides shortcuts for common commands:
```Makefile
run:
	go run cmd/main.go

migrate:
	go run internal/db/migrate.go

docker-build:
	docker build -t go-backend-api .

docker-run:
	docker run -p 8080:8080 --env-file .env go-backend-api
```

### 🔹 Available Commands
```sh
make run         # Start the Go application
make migrate     # Run database migrations
make docker-build # Build the Docker image
make docker-run   # Run the containerized application
```

---

## 🔥 API Endpoints

| Method | Endpoint         | Description         |
|--------|-----------------|---------------------|
| POST   | /register       | Register a user    |
| POST   | /login          | Authenticate user  |
| GET    | /users          | Get all users      |
| GET    | /users/{id}     | Get user by ID     |

---

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

🚀 **Happy Coding!**

