# Go Interview API

A RESTful API built with Go and SQLite for serving and managing curated Go interview questions. Designed for learning, contribution, and portfolio demonstration.

---

## ✨ Features

- 📚 View Go interview questions
- 🔍 Filter by topic
- 🔐 Admin routes (create, edit, delete questions)
- 🪪 JWT-based authentication
- 💬 Markdown-to-HTML support
- 🧩 Swagger/OpenAPI documentation (optional)
- 📁 Clean folder structure (ready for contributions)

---

## ⚙️ Tech Stack

- Go (Golang)
- Chi Router
- GORM
- SQLite (via modernc.org or glebarez/sqlite)
- JWT for authentication

---

## 🚀 Getting Started

Clone the repository

    git clone https://github.com/yourusername/go-interview-api.git
    cd go-interview-api

Install dependencies

    go mod tidy

Run the API

    go run cmd/server/main.go

---

## 📦 API Endpoints

### 🔐 Auth

| Method | Route             | Description           |
|--------|-------------------|-----------------------|
| POST   | /users/register   | Register new user     |
| POST   | /users/login      | Login and get token   |

### 📄 Questions (Protected)

| Method | Route             | Description             |
|--------|-------------------|-------------------------|
| GET    | /questions        | List all questions      |
| POST   | /questions        | Create a new question   |

Add your JWT in requests:
    
    Authorization: Bearer <token>

---

## 🧪 Testing

You can test with:
- Postman
- Curl
- Swagger UI (coming soon)

---

## 🤝 Contributing

Contributions are welcome! Please open an issue or submit a pull request.

---

## 📄 License

MIT License

---

## 👤 Author

Made with ❤️ by @Fred-CodeCrafts
