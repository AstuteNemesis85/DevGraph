# 🚀 DevGraph - Complete Setup Guide

## Backend Setup (Already Complete)

Your Go backend is ready! Just need to install CORS package:

```bash
cd d:\CodeScope\devgraph
go get github.com/gin-contrib/cors
go mod tidy
```

Then start the backend:
```bash
go run cmd/server/main.go
```

The backend will run on `http://localhost:8080`

---

## Frontend Setup (NEW)

### 1. Install Dependencies

```bash
cd d:\CodeScope\devgraph\frontend
npm install
```

### 2. Start Development Server

```bash
npm run dev
```

The frontend will open at `http://localhost:3000`

---

## 🎯 How to Use DevGraph

### 1. **Register an Account**
- Navigate to `http://localhost:3000`
- Click "Sign up" 
- Enter username, email, and password (min 8 chars)

### 2. **Login**
- Use your email and password
- You'll be redirected to the Dashboard

### 3. **Submit Code**
- Select a programming language (Python, JavaScript, Java, C++, Go)
- Write or paste your code in the Monaco editor
- Click "Submit for Analysis"
- Your code will be analyzed in the background

### 4. **View Recommendations**
- Click "Recommendations" in the navbar
- See developers with similar coding patterns
- View similarity scores and shared algorithm patterns

---

## 🎨 UI Features

### ✨ Modern Design
- Gradient backgrounds and cards
- Smooth animations and transitions
- Professional color scheme (Blue & Purple)
- Responsive layout for all screen sizes

### 💻 Code Editor
- **Monaco Editor** (same as VS Code)
- Syntax highlighting
- Auto-completion
- Line numbers
- Dark theme

### 🔐 Authentication
- Secure JWT tokens
- Automatic token refresh
- Protected routes
- Beautiful login/register pages

### 📊 Dashboard
- Real-time code submissions
- Submission history sidebar
- Language selector
- Status indicators

### 👥 Recommendations Page
- Beautiful developer cards
- Similarity badges (Very High, High, Medium, Low)
- Shared pattern counts
- Color-coded similarity scores

---

## 🛠️ Tech Stack

### Backend (Go)
- Gin web framework
- PostgreSQL database
- Redis caching
- JWT authentication
- Background workers for analysis

### Frontend (React)
- React 18 with hooks
- Vite for fast development
- Tailwind CSS for styling
- Monaco Editor for code editing
- Axios for API calls
- React Router for navigation
- React Hot Toast for notifications

---

## 📁 Project Structure

```
devgraph/
├── backend (Go)
│   ├── cmd/server/main.go
│   ├── internal/
│   │   ├── auth/
│   │   ├── code/
│   │   ├── analysis/
│   │   ├── graph/
│   │   └── user/
│   └── go.mod
│
└── frontend/ (NEW - React)
    ├── src/
    │   ├── components/
    │   ├── context/
    │   ├── pages/
    │   ├── services/
    │   └── App.jsx
    ├── package.json
    └── vite.config.js
```

---

## 🔧 Configuration

### Backend (.env)
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=devgraph
REDIS_ADDR=localhost:6379
JWT_SECRET=your_secret_key
```

### Frontend (Built-in Proxy)
The Vite config automatically proxies API calls to `http://localhost:8080`

---

## 🚨 Troubleshooting

### CORS Errors
✅ Already fixed! CORS middleware is configured in main.go

### Port Already in Use
```bash
# Backend (if port 8080 is taken)
netstat -ano | findstr :8080
taskkill /PID <pid> /F

# Frontend (if port 3000 is taken)
netstat -ano | findstr :3000
taskkill /PID <pid> /F
```

### Monaco Editor Not Loading
- Clear browser cache
- Restart the development server
- Check browser console for errors

### API Connection Failed
- Ensure backend is running on port 8080
- Check CORS configuration
- Verify .env file has correct values

---

## 🎉 What You Get

### 1. **Beautiful Landing Experience**
- Modern gradient design
- Professional branding
- Smooth animations

### 2. **Powerful Code Editor**
- VS Code-quality editor
- Multi-language support
- Real-time syntax highlighting

### 3. **Smart Analysis**
- Background processing
- Pattern detection
- Complexity analysis

### 4. **Social Features**
- Developer recommendations
- Similarity matching
- Shared coding patterns

---

## 📸 Color Palette

- **Primary Blue**: `#0ea5e9` to `#0284c7`
- **Primary Purple**: `#9333ea` to `#7e22ce`
- **Success**: `#10b981`
- **Warning**: `#f59e0b`
- **Error**: `#ef4444`

---

## 🚀 Production Deployment

### Frontend
```bash
npm run build
# Deploy the `dist` folder to Netlify, Vercel, or any static host
```

### Backend
```bash
go build -o devgraph cmd/server/main.go
# Deploy to AWS, GCP, or any Go hosting service
```

---

## 📝 Next Steps

1. ✅ Start both backend and frontend
2. ✅ Register a new account
3. ✅ Submit some code samples
4. ✅ Check out the recommendations
5. 🎨 Customize colors in `tailwind.config.js`
6. 🔧 Add more features as needed

---

## 🎓 Learning Resources

- [React Docs](https://react.dev)
- [Tailwind CSS](https://tailwindcss.com)
- [Monaco Editor](https://microsoft.github.io/monaco-editor/)
- [Gin Framework](https://gin-gonic.com)

---

**Enjoy your modern DevGraph application! 🎉**
