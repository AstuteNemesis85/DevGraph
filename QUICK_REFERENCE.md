# 🚀 DevGraph - Quick Reference Card

## 📍 URLs
- **Backend API**: http://localhost:8080
- **Frontend App**: http://localhost:3000

---

## ⚡ Quick Start Commands

### Start Everything (Windows)
```bash
# Terminal 1 - Backend
cd d:\CodeScope\devgraph
go get github.com/gin-contrib/cors
./start-backend.bat

# Terminal 2 - Frontend
cd d:\CodeScope\devgraph
./start-frontend.bat
```

### Start Everything (Linux/Mac)
```bash
# Terminal 1 - Backend
cd ~/CodeScope/devgraph
go get github.com/gin-contrib/cors
chmod +x start-backend.sh
./start-backend.sh

# Terminal 2 - Frontend
cd ~/CodeScope/devgraph
chmod +x start-frontend.sh
./start-frontend.sh
```

---

## 📁 Important Files

| File | Purpose |
|------|---------|
| `cmd/server/main.go` | Backend entry point |
| `frontend/src/App.jsx` | Frontend app & routing |
| `frontend/src/services/api.js` | API client |
| `SETUP_GUIDE.md` | Detailed setup |
| `API_DOCUMENTATION.md` | API reference |

---

## 🎨 Pages

| Route | Page | Description |
|-------|------|-------------|
| `/` | Redirect | → `/dashboard` |
| `/login` | Login | Email + password auth |
| `/register` | Register | Create new account |
| `/dashboard` | Dashboard | Code editor + submissions |
| `/recommendations` | Recommendations | Similar developers |

---

## 🔌 API Endpoints

| Method | Endpoint | Auth | Purpose |
|--------|----------|------|---------|
| POST | `/auth/register` | ❌ | Register account |
| POST | `/auth/login` | ❌ | Login |
| POST | `/auth/refresh` | ❌ | Refresh token |
| POST | `/auth/logout` | ✅ | Logout |
| GET | `/api/me` | ✅ | Get user info |
| POST | `/api/submit` | ✅ | Submit code |
| GET | `/api/recommendations` | ✅ | Get similar devs |

---

## 🛠️ Tech Stack

### Backend
- Go 1.24
- Gin (web framework)
- PostgreSQL (database)
- Redis (cache)
- JWT (auth)

### Frontend
- React 18
- Vite (bundler)
- Tailwind CSS
- Monaco Editor
- Axios

---

## 🎨 Color Palette

```css
Primary Blue: #0ea5e9
Primary Purple: #9333ea
Success: #10b981
Warning: #f59e0b
Error: #ef4444
```

---

## 📦 Languages Supported

- Python
- JavaScript
- Java
- C++
- Go

---

## 🔑 Environment Variables

```env
# Required in .env file
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=devgraph
REDIS_ADDR=localhost:6379
JWT_SECRET=your_secret_key
```

---

## 🐛 Troubleshooting

### Port Already in Use
```bash
# Windows
netstat -ano | findstr :8080
taskkill /PID <pid> /F

# Linux/Mac
lsof -ti:8080 | xargs kill -9
```

### CORS Error
✅ Already fixed! CORS middleware added to main.go

### Frontend Won't Start
```bash
cd frontend
rm -rf node_modules
npm install
npm run dev
```

### Backend Won't Start
```bash
go mod tidy
go get github.com/gin-contrib/cors
go run cmd/server/main.go
```

---

## 📚 Documentation Links

- [SETUP_GUIDE.md](./SETUP_GUIDE.md) - Step-by-step setup
- [API_DOCUMENTATION.md](./API_DOCUMENTATION.md) - Complete API docs
- [FRONTEND_SUMMARY.md](./FRONTEND_SUMMARY.md) - Frontend overview
- [frontend/README.md](./frontend/README.md) - Frontend details

---

## 🎯 Common Tasks

### Register New User
1. Go to http://localhost:3000
2. Click "Sign up"
3. Enter username, email, password
4. Click "Create Account"

### Submit Code
1. Login to dashboard
2. Select language
3. Write code in editor
4. Click "Submit for Analysis"

### View Recommendations
1. Click "Recommendations" in navbar
2. See similar developers
3. View similarity scores

---

## 🔒 Security Notes

- Tokens stored in localStorage
- Access token: 15 min expiry
- Refresh token: 7 day expiry
- Automatic token refresh
- Password hashed with bcrypt

---

## 📈 Future Ideas

- [ ] User profiles
- [ ] Code history with details
- [ ] Analysis visualization
- [ ] Dark mode
- [ ] Real-time notifications
- [ ] Social features
- [ ] Leaderboards

---

## ✨ Quick Tips

1. Keep both servers running
2. Check browser console for errors
3. Backend logs show requests
4. Clear cache if issues occur
5. Use React DevTools for debugging

---

## 🆘 Need Help?

1. Check [SETUP_GUIDE.md](./SETUP_GUIDE.md)
2. Check [API_DOCUMENTATION.md](./API_DOCUMENTATION.md)
3. Check browser console
4. Check backend terminal logs
5. Check .env configuration

---

**You're all set! Start coding! 🚀**
