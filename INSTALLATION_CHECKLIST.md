# ✅ DevGraph Installation Checklist

## Prerequisites Check

- [ ] Go 1.24+ installed (`go version`)
- [ ] Node.js 18+ installed (`node --version`)
- [ ] PostgreSQL installed and running
- [ ] Redis installed and running
- [ ] Git installed (optional)

---

## Backend Setup

### 1. Install Dependencies
```bash
cd d:\CodeScope\devgraph
go mod tidy
go get github.com/gin-contrib/cors
```
- [ ] Dependencies installed successfully

### 2. Configure Environment
```bash
# Create .env file in project root
cp .env.example .env
# Edit .env with your database credentials
```
- [ ] .env file created
- [ ] Database credentials configured
- [ ] Redis address configured
- [ ] JWT secret set

### 3. Database Setup
- [ ] PostgreSQL running on configured port
- [ ] Database created (or will be auto-created)
- [ ] Connection test successful

### 4. Redis Setup
- [ ] Redis running on configured port
- [ ] Redis accessible

### 5. Start Backend
```bash
./start-backend.bat  # Windows
# or
./start-backend.sh   # Linux/Mac
```
- [ ] Backend starts without errors
- [ ] Server running on http://localhost:8080
- [ ] Database migrations complete
- [ ] Worker pool started

---

## Frontend Setup

### 1. Install Dependencies
```bash
cd d:\CodeScope\devgraph\frontend
npm install
```
- [ ] Node modules installed (~2-3 minutes)
- [ ] No dependency errors

### 2. Verify Configuration
- [ ] `vite.config.js` proxy set to `http://localhost:8080`
- [ ] `tailwind.config.js` exists
- [ ] `package.json` has all dependencies

### 3. Start Frontend
```bash
cd d:\CodeScope\devgraph
./start-frontend.bat  # Windows
# or
./start-frontend.sh   # Linux/Mac
```
- [ ] Frontend starts without errors
- [ ] Vite dev server running
- [ ] Opens browser at http://localhost:3000

---

## Verification Tests

### Backend Tests
```bash
# Test health endpoint
curl http://localhost:8080/auth/login
# Should return "email" and "password" required error
```
- [ ] Backend responds to requests
- [ ] CORS headers present

### Frontend Tests
- [ ] http://localhost:3000 loads
- [ ] Login page displays correctly
- [ ] No console errors in browser
- [ ] Images and styles load

### Integration Tests
- [ ] Register new user works
- [ ] Login works
- [ ] Redirects to dashboard after login
- [ ] Monaco editor loads
- [ ] Can submit code
- [ ] Can view recommendations

---

## First-Time User Flow

### 1. Register Account
- [ ] Go to http://localhost:3000
- [ ] Click "Sign up"
- [ ] Enter username: `testuser`
- [ ] Enter email: `test@example.com`
- [ ] Enter password: `password123`
- [ ] Click "Create Account"
- [ ] Success message appears
- [ ] Redirected to login

### 2. Login
- [ ] Enter email: `test@example.com`
- [ ] Enter password: `password123`
- [ ] Click "Sign In"
- [ ] Redirected to dashboard

### 3. Submit Code
- [ ] Monaco editor visible
- [ ] Select language: Python
- [ ] Write simple code:
```python
def hello():
    print("Hello, World!")
```
- [ ] Click "Submit for Analysis"
- [ ] Success toast appears
- [ ] Submission appears in sidebar

### 4. View Recommendations
- [ ] Click "Recommendations" in navbar
- [ ] Page loads (may be empty at first)
- [ ] No errors in console

---

## Common Issues & Fixes

### Issue: Port 8080 already in use
```bash
# Windows
netstat -ano | findstr :8080
taskkill /PID <pid> /F

# Linux/Mac
lsof -ti:8080 | xargs kill -9
```
- [ ] Port freed
- [ ] Backend restarted

### Issue: Port 3000 already in use
```bash
# Windows
netstat -ano | findstr :3000
taskkill /PID <pid> /F

# Linux/Mac
lsof -ti:3000 | xargs kill -9
```
- [ ] Port freed
- [ ] Frontend restarted

### Issue: Database connection failed
- [ ] PostgreSQL running
- [ ] Credentials correct in .env
- [ ] Database name exists
- [ ] Port correct (default 5432)

### Issue: Redis connection failed
- [ ] Redis running
- [ ] Address correct in .env (default localhost:6379)
- [ ] Redis accessible

### Issue: CORS error in browser
- [ ] Backend has CORS middleware
- [ ] Backend restarted after adding CORS
- [ ] Browser cache cleared

### Issue: Monaco editor not loading
- [ ] Internet connection active (for CDN)
- [ ] Browser cache cleared
- [ ] Dev server restarted

### Issue: npm install fails
```bash
# Clear npm cache
npm cache clean --force
rm -rf node_modules package-lock.json
npm install
```
- [ ] Dependencies reinstalled

---

## Performance Check

### Backend
- [ ] Response time < 100ms for auth endpoints
- [ ] Database queries fast
- [ ] Worker pool processing submissions
- [ ] Redis caching working

### Frontend
- [ ] Page load < 2 seconds
- [ ] Smooth animations
- [ ] Editor responsive
- [ ] No memory leaks

---

## Documentation Review

- [ ] Read [README.md](./README.md)
- [ ] Read [SETUP_GUIDE.md](./SETUP_GUIDE.md)
- [ ] Skim [API_DOCUMENTATION.md](./API_DOCUMENTATION.md)
- [ ] Check [QUICK_REFERENCE.md](./QUICK_REFERENCE.md)
- [ ] Review [frontend/README.md](./frontend/README.md)

---

## Optional Setup

### Git Configuration
```bash
git init
git add .
git commit -m "Initial commit with frontend"
```
- [ ] Git repository initialized
- [ ] Initial commit made

### Production Build
```bash
# Frontend production build
cd frontend
npm run build
# Creates optimized build in dist/
```
- [ ] Production build successful
- [ ] Build size reasonable

### Database Backup
```bash
# Create backup script
pg_dump devgraph > backup.sql
```
- [ ] Backup working

---

## Final Checklist

### Backend ✅
- [ ] Installed and configured
- [ ] Running on port 8080
- [ ] Database connected
- [ ] Redis connected
- [ ] CORS enabled
- [ ] JWT working
- [ ] Worker pool active

### Frontend ✅
- [ ] Installed and configured
- [ ] Running on port 3000
- [ ] API proxy working
- [ ] All pages load
- [ ] Monaco editor works
- [ ] Authentication works
- [ ] No console errors

### Testing ✅
- [ ] Registration works
- [ ] Login works
- [ ] Code submission works
- [ ] Navigation works
- [ ] Logout works
- [ ] Token refresh works

### Documentation ✅
- [ ] All docs created
- [ ] README updated
- [ ] API docs complete
- [ ] Setup guide clear

---

## 🎉 Installation Complete!

If all checkboxes are checked, your DevGraph installation is complete and ready for development!

### Next Steps:
1. Start coding with the platform
2. Submit multiple code samples
3. Wait for analysis to complete
4. Check recommendations
5. Customize styling if needed
6. Add new features

---

## 🆘 Still Having Issues?

1. Check backend terminal for error logs
2. Check browser console for errors
3. Verify .env file configuration
4. Ensure all services running (PostgreSQL, Redis)
5. Try restarting both servers
6. Clear browser cache and localStorage
7. Check firewall settings

---

## 📞 Support Resources

- Backend logs: Terminal running backend
- Frontend logs: Browser DevTools console
- Network logs: Browser DevTools Network tab
- Database: pgAdmin or psql
- Redis: redis-cli

---

**Installation checklist complete! Happy coding! 🚀**

