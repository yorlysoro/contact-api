---

# Contact Management API (Golang)

A robust, production-ready RESTful API for managing contacts with hierarchical relationships (Parent/Child). Built using **Clean Architecture**, **TDD**, and the **Gin/GORM** ecosystem.

## 🚀 Features

* **Full CRUD:** Create, Read, Update, and Delete contacts.
* **Hierarchical Contacts:** Support for parent-child contact relationships using GORM self-referencing.
* **JWT Authentication:** Secure endpoints using JSON Web Tokens.
* **Clean Architecture:** Strict separation of concerns (Handler -> Service -> Repository).
* **TDD Approach:** High test coverage using the `testify` framework and mocks.
* **Relational Database:** Ready for SQLite (default) or PostgreSQL/MySQL.

## 🛠️ Tech Stack

* **Language:** Go 1.21+
* **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
* **ORM:** [GORM](https://gorm.io/)
* **Security:** [JWT-Go](https://github.com/golang-jwt/jwt)
* **Testing:** [Testify](https://github.com/stretchr/testify)

---

## 📂 Project Structure

```text
contact-api/
├── cmd/                # Application entry point (main.go)
├── internal/           
│   ├── auth/           # JWT Middleware and token logic
│   ├── contact/        # Contact domain logic (Handler, Service, Repo)
│   └── models/         # GORM Entities (Database schemas)
├── pkg/                # Shared utilities (DB connections, etc.)
├── tests/              # Integration and unit tests
└── go.mod              # Dependency management

```

---

## ⚙️ Installation & Setup

### 1. Prerequisites

* [Go](https://go.dev/doc/install) installed.
* A terminal/command prompt.

### 2. Clone and Initialize

```bash
git clone https://github.com/youruser/contact-api.git
cd contact-api
go mod tidy

```

### 3. Environment Configuration

The project uses a secret key for JWT. For local development, you can export it:

```bash
export JWT_SECRET="your-super-secret-key"

```

---

## 🏃 Running the Project

### Start the Server

```bash
go run cmd/main.go

```

The server will start at `http://localhost:8080`.

### Run Tests (TDD)

To verify the business logic and API endpoints:

```bash
go test ./... -v

```

---

## 📡 API Endpoints

| Method | Endpoint | Description | Auth Required |
| --- | --- | --- | --- |
| **POST** | `/api/v1/contacts` | Create a new contact | Yes (JWT) |
| **GET** | `/api/v1/contacts/:id` | Get contact with children | Yes (JWT) |
| **PUT** | `/api/v1/contacts/:id` | Update contact | Yes (JWT) |
| **DELETE** | `/api/v1/contacts/:id` | Delete contact | Yes (JWT) |

---

## 📝 Usage Example: Creating a Related Contact

**Request:**
`POST /api/v1/contacts`

```json
{
    "name": "Jane Doe",
    "phone": "555-0199",
    "parent_id": 1
}

```

**Response:**
`201 Created`

```json
{
    "id": 2,
    "name": "Jane Doe",
    "phone": "555-0199",
    "parent_id": 1,
    "user_id": 10
}

```

---

## 🤝 Contributing

1. Create a Feature Branch.
2. Ensure all tests pass (`go test ./...`).
3. Submit a Pull Request.

---
