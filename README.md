# DevGraph 🚀

**A modern code analysis and developer networking platform**

DevGraph helps developers submit code for automated analysis and discover other developers with similar coding patterns and interests.

![Tech Stack](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![React](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white)

---

## ✨ Features

- 🔐 **Secure Authentication** - JWT-based auth with automatic token refresh
- 💻 **Code Editor** - Monaco Editor (VS Code) with multi-language support
- 📊 **Code Analysis** - Automatic complexity and pattern detection
- 👥 **Developer Recommendations** - Find similar developers based on coding patterns
- 🎨 **Modern UI** - Beautiful, responsive interface with Tailwind CSS
- ⚡ **Real-time Updates** - Live submission tracking and notifications

---

## 🚀 Quick Start

### Prerequisites
- Go 1.24+
- Node.js 18+
- PostgreSQL
- Redis

### Backend Setup
```bash
# Install dependencies
go mod tidy
go get github.com/gin-contrib/cors

# Start backend (port 8080)
./start-backend.bat  # Windows
./start-backend.sh   # Linux/Mac
```

### Frontend Setup
```bash
# Install and start (port 3000)
./start-frontend.bat  # Windows
./start-frontend.sh   # Linux/Mac
```

---

## 📁 Project Structure

```
devgraph/
├── cmd/server/main.go              # Entry point
├── internal/
│   ├── auth/                       # Authentication
│   ├── code/                       # Code submission
│   ├── analysis/                   # Analysis engine
│   ├── graph/                      # Recommendations
│   ├── user/                       # User management
│   ├── cache/                      # Redis cache
│   └── config/                     # Database
├── frontend/                       # React app
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   └── services/
│   └── package.json
├── SETUP_GUIDE.md                  # Detailed guide
└── API_DOCUMENTATION.md            # API reference
```

---

## 🔌 API Endpoints

- `POST /auth/register` - Register
- `POST /auth/login` - Login
- `POST /api/submit` - Submit code
- `GET /api/recommendations` - Get similar devs

See [API_DOCUMENTATION.md](./API_DOCUMENTATION.md) for details.

---

## 🛠️ Tech Stack

**Backend:** Go, Gin, PostgreSQL, Redis, JWT  
**Frontend:** React, Vite, Tailwind, Monaco Editor

---

## 📚 Documentation

- [SETUP_GUIDE.md](./SETUP_GUIDE.md) - Complete setup
- [API_DOCUMENTATION.md](./API_DOCUMENTATION.md) - API docs
- [frontend/README.md](./frontend/README.md) - Frontend docs

---

## Getting Started

1. Copy `.env.example` to `.env` and configure
2. Start backend: `./start-backend.bat`
3. Start frontend: `./start-frontend.bat`
4. Open `http://localhost:3000`

---


