# 🚀 DevGraph - Deployment Guide

This guide provides step-by-step instructions for deploying DevGraph in different environments.

---

## 📋 Table of Contents
- [⚡ Quick Deploy: Render + Vercel (Recommended)](#-quick-deploy-render--vercel-recommended)
- [Development Deployment (Local)](#development-deployment-local)
- [Production Deployment (VPS/Cloud)](#production-deployment-vpscloud)
- [Docker Deployment](#docker-deployment)
- [Platform-Specific Guides](#platform-specific-guides)

---

## ⚡ Quick Deploy: Render + Vercel (Recommended)

**Best for:** Quick production deployment with minimal configuration, free tier available

This is the **simplest production deployment method** - no server management required!

- **Backend + Database + Redis:** Render (all in one platform)
- **Frontend:** Vercel (lightning-fast CDN)
- **Time:** ~15 minutes
- **Cost:** Free tier available (or ~$7-25/month for production)

### Prerequisites
- GitHub account
- Render account ([render.com](https://render.com))
- Vercel account ([vercel.com](https://vercel.com))
- Your code pushed to GitHub

---

### Part A: Deploy Backend on Render

#### Step 1: Push Code to GitHub

```bash
cd d:\CodeScope\devgraph
git init
git add .
git commit -m "Initial commit"
git remote add origin https://github.com/yourusername/devgraph.git
git push -u origin main
```

#### Step 2: Create PostgreSQL Database

1. Go to [Render Dashboard](https://dashboard.render.com)
2. Click **"New +"** → **"PostgreSQL"**
3. Configure:
   - **Name:** `devgraph-db`
   - **Database:** `devgraph`
   - **User:** `devgraph_user`
   - **Region:** Choose closest to your users
   - **Plan:** Free or Starter ($7/month)
4. Click **"Create Database"**
5. **Save the Internal Database URL** (you'll need this)

#### Step 3: Create Redis Instance

1. Click **"New +"** → **"Redis"**
2. Configure:
   - **Name:** `devgraph-redis`
   - **Region:** Same as database
   - **Plan:** Free or Starter ($3/month)
3. Click **"Create Redis"**
4. **Save the Internal Redis URL** (you'll need this)

#### Step 4: Deploy Backend Web Service

1. Click **"New +"** → **"Web Service"**
2. Connect your GitHub repository
3. Configure:
   - **Name:** `devgraph-backend`
   - **Region:** Same as database
   - **Branch:** `main`
   - **Root Directory:** Leave empty (root)
   - **Runtime:** `Go`
   - **Build Command:**
     ```bash
     go build -o devgraph-server cmd/server/main.go
     ```
   - **Start Command:**
     ```bash
     ./devgraph-server
     ```
   - **Plan:** Free or Starter ($7/month)

4. Click **"Advanced"** and add environment variables:

   ```env
   DB_HOST=<from-postgres-internal-connection>
   DB_PORT=5432
   DB_USER=devgraph_user
   DB_PASSWORD=<from-postgres-connection>
   DB_NAME=devgraph
   
   JWT_SECRET=<generate-random-32-char-string>
   
   SERVER_PORT=8080
   
   REDIS_URL=redis://red-d6bvgh9r8fns73ar9d20.internal:6379
   
   GIN_MODE=release
   ```

   **To get connection details:**
   - Click on your PostgreSQL database → "Info" tab → Copy "Internal Database URL"
   - Click on your Redis instance → "Info" tab → Copy "Internal Redis URL" (use the full URL)
   - For JWT_SECRET, generate with: `openssl rand -base64 32`

5. Click **"Create Web Service"**

6. **Wait for deployment** (5-10 minutes first time)

7. Once deployed, **copy your backend URL**: `https://devgraph-backend.onrender.com`

#### Step 5: Update CORS Settings

Before deploying frontend, update the CORS configuration in your backend:

```bash
# Edit cmd/server/main.go locally
```

Find the CORS section and update:

```go
r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{
        "http://localhost:3000",
        "https://devgraph.vercel.app",           // Your Vercel URL
        "https://your-custom-domain.com",        // If you have one
    },
    AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
}))
```

```bash
# Commit and push changes
git add cmd/server/main.go
git commit -m "Update CORS for production"
git push
```

Render will auto-deploy the changes.

---

### Part B: Deploy Frontend on Vercel

#### Step 1: Update Frontend API Configuration

Edit `frontend/src/services/api.js`:

```javascript
// Replace localhost with your Render backend URL
const API_BASE_URL = import.meta.env.PROD 
  ? 'https://devgraph-backend.onrender.com'  // Your Render URL
  : 'http://localhost:8080';

export default API_BASE_URL;
```

Commit changes:
```bash
git add frontend/src/services/api.js
git commit -m "Configure production API URL"
git push
```

#### Step 2: Deploy to Vercel

1. Go to [Vercel Dashboard](https://vercel.com/dashboard)
2. Click **"Add New..."** → **"Project"**
3. **Import** your GitHub repository
4. Configure:
   - **Framework Preset:** Vite
   - **Root Directory:** `frontend`
   - **Build Command:** `npm run build`
   - **Output Directory:** `dist`
   - **Install Command:** `npm install`

5. Click **"Deploy"**

6. **Wait for deployment** (2-3 minutes)

7. Your app will be live at: `https://devgraph.vercel.app`

#### Step 3: Configure Custom Domain (Optional)

1. In Vercel project settings → **"Domains"**
2. Add your custom domain
3. Update DNS records as instructed
4. SSL certificate is automatically provisioned

---

### Part C: Testing Your Deployment

1. Visit your Vercel URL: `https://devgraph.vercel.app`
2. Register a new account
3. Login
4. Submit code for analysis
5. Check recommendations

---

### 🎯 Deployment Summary

| Component | Platform | URL | Cost |
|-----------|----------|-----|------|
| Frontend | Vercel | `https://devgraph.vercel.app` | Free |
| Backend | Render | `https://devgraph-backend.onrender.com` | $7/month |
| PostgreSQL | Render | Internal | $7/month |
| Redis | Render | Internal | $3/month |
| **Total** | | | **$17/month** or Free tier |

---

### 🔄 Auto-Deploy Updates

Both platforms support automatic deployments:

**Vercel:**
- Automatically redeploys on every `git push` to `main`
- Preview deployments for pull requests

**Render:**
- Automatically redeploys on every `git push` to `main`
- Can configure manual deploy if preferred

To update your app:
```bash
# Make changes locally
git add .
git commit -m "Your update message"
git push

# Both platforms auto-deploy in minutes!
```

---

### 💡 Pro Tips for Render + Vercel

**1. Free Tier Considerations:**
- Render free tier: Backend sleeps after 15 min inactivity (cold start ~30s)
- Vercel free tier: Perfect for most use cases
- Upgrade to paid for production traffic

**2. Environment Variables:**
- Update in Render dashboard under "Environment" tab
- Changes require manual redeploy

**3. Monitoring:**
- Render: Built-in logs and metrics
- Vercel: Built-in analytics and deployment logs

**4. Database Backups:**
- Render PostgreSQL: Automatic daily backups (paid plans)
- Manual backup: Download from Render dashboard

**5. Performance:**
- Use same region for all Render services
- Enable Vercel's Edge Network

**6. Debugging:**
- Render logs: Dashboard → Your service → "Logs" tab
- Vercel logs: Dashboard → Your project → "Deployments" → Click deployment
- Check backend health: `https://devgraph-backend.onrender.com/api/health` (add health endpoint)

---

### ❗ Common Issues & Fixes

**Issue: Backend won't start**
```bash
# Check Render logs for errors
# Common fix: Verify all environment variables are set
# Go to Render service → Environment → Add missing vars
```

**Issue: Frontend can't connect to backend**
```bash
# 1. Check CORS settings in cmd/server/main.go
# 2. Verify API_BASE_URL in frontend/src/services/api.js
# 3. Check browser console for CORS errors
```

**Issue: Database connection failed**
```bash
# Verify DB environment variables in Render
# Use Internal Database URL, not External
# Format: host=xxx port=5432 user=xxx password=xxx dbname=xxx
```

**Issue: Cold start delays (free tier)**
```bash
# Upgrade to paid Render plan ($7/month)
# Or: Accept 30s first-request delay on free tier
```

---

### 🚀 Next Steps

After deployment:
- [ ] Configure custom domain
- [ ] Set up monitoring/alerts
- [ ] Enable auto-scaling (paid plans)
- [ ] Configure backup strategy
- [ ] Set up staging environment
- [ ] Add health check endpoints
- [ ] Configure rate limiting
- [ ] Set up error tracking (Sentry)

---

## 🔧 Development Deployment (Local)

### Prerequisites
- **Go** 1.24+ ([Download](https://golang.org/dl/))
- **Node.js** 18+ ([Download](https://nodejs.org/))
- **PostgreSQL** 14+ ([Download](https://www.postgresql.org/download/))
- **Redis** 7+ ([Download](https://redis.io/download))
- **Git**

### Step 1: Clone and Setup Repository

```bash
# Clone the repository
git clone <your-repo-url>
cd devgraph
```

### Step 2: Setup PostgreSQL Database

```bash
# Start PostgreSQL service
# Windows: Services -> PostgreSQL -> Start
# Linux: sudo systemctl start postgresql
# Mac: brew services start postgresql

# Create database
psql -U postgres
CREATE DATABASE devgraph;
CREATE USER devgraph_user WITH PASSWORD 'your_secure_password';
GRANT ALL PRIVILEGES ON DATABASE devgraph TO devgraph_user;
\q
```

### Step 3: Setup Redis

```bash
# Windows: Download and run redis-server.exe
# Linux: sudo systemctl start redis
# Mac: brew services start redis

# Verify Redis is running
redis-cli ping
# Should return: PONG
```

### Step 4: Configure Environment Variables

```bash
# Copy example env file
cp .env.example .env

# Edit .env with your credentials
```

**Sample .env file:**
```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=devgraph_user
DB_PASSWORD=your_secure_password
DB_NAME=devgraph

# JWT Secret (generate with: openssl rand -base64 32)
JWT_SECRET=your_jwt_secret_key_here_make_it_long_and_random

# Server Configuration
SERVER_PORT=8080

# Redis Configuration (if using custom settings)
REDIS_HOST=localhost
REDIS_PORT=6379
```

### Step 5: Install Backend Dependencies

```bash
# Install Go modules
go mod download
go mod tidy
```

### Step 6: Run Database Migrations

The application automatically runs migrations on startup, but you can verify:

```bash
# Start backend once to run migrations
go run cmd/server/main.go
# Stop after you see "Server running on :8080"
```

### Step 7: Start Backend Server

```bash
# Option 1: Using script
./start-backend.bat  # Windows
./start-backend.sh   # Linux/Mac

# Option 2: Direct command
go run cmd/server/main.go
```

Backend should be running at `http://localhost:8080`

### Step 8: Setup Frontend

```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

Frontend should be running at `http://localhost:3000`

### Step 9: Verify Installation

1. Open browser to `http://localhost:3000`
2. Register a new account
3. Login and test code submission
4. Check recommendations page

---

## 🌐 Production Deployment (VPS/Cloud)

### Prerequisites
- VPS or cloud instance (AWS EC2, DigitalOcean, Linode, etc.)
- Ubuntu 22.04 LTS (or similar)
- Domain name (optional but recommended)
- SSL certificate (Let's Encrypt recommended)

### Step 1: Server Setup

```bash
# Connect to your server
ssh user@your-server-ip

# Update system
sudo apt update && sudo apt upgrade -y

# Install dependencies
sudo apt install -y build-essential git nginx certbot python3-certbot-nginx
```

### Step 2: Install Go

```bash
# Download and install Go
cd /tmp
wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz

# Add to PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verify installation
go version
```

### Step 3: Install Node.js

```bash
# Install Node.js 18.x
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt install -y nodejs

# Verify installation
node --version
npm --version
```

### Step 4: Install PostgreSQL

```bash
# Install PostgreSQL
sudo apt install -y postgresql postgresql-contrib

# Start and enable PostgreSQL
sudo systemctl start postgresql
sudo systemctl enable postgresql

# Create database and user
sudo -u postgres psql
CREATE DATABASE devgraph;
CREATE USER devgraph_user WITH PASSWORD 'STRONG_PRODUCTION_PASSWORD';
GRANT ALL PRIVILEGES ON DATABASE devgraph TO devgraph_user;
\q
```

### Step 5: Install Redis

```bash
# Install Redis
sudo apt install -y redis-server

# Configure Redis for production
sudo nano /etc/redis/redis.conf
# Set: supervised systemd
# Set: bind 127.0.0.1

# Restart Redis
sudo systemctl restart redis
sudo systemctl enable redis

# Verify
redis-cli ping
```

### Step 6: Clone and Build Application

```bash
# Create application directory
sudo mkdir -p /opt/devgraph
sudo chown $USER:$USER /opt/devgraph
cd /opt/devgraph

# Clone repository
git clone <your-repo-url> .

# Create and configure .env
nano .env
```

**Production .env:**
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=devgraph_user
DB_PASSWORD=STRONG_PRODUCTION_PASSWORD
DB_NAME=devgraph

JWT_SECRET=YOUR_VERY_LONG_RANDOM_JWT_SECRET_HERE

SERVER_PORT=8080

REDIS_HOST=localhost
REDIS_PORT=6379

# Production mode
GIN_MODE=release
```

```bash
# Build backend
go mod download
go build -o devgraph-server cmd/server/main.go

# Build frontend
cd frontend
npm install
npm run build
cd ..
```

### Step 7: Setup Systemd Service for Backend

```bash
# Create systemd service
sudo nano /etc/systemd/system/devgraph.service
```

**Service file content:**
```ini
[Unit]
Description=DevGraph Backend Service
After=network.target postgresql.service redis.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/devgraph
ExecStart=/opt/devgraph/devgraph-server
Restart=always
RestartSec=10

Environment="GIN_MODE=release"

[Install]
WantedBy=multi-user.target
```

```bash
# Reload systemd and start service
sudo systemctl daemon-reload
sudo systemctl start devgraph
sudo systemctl enable devgraph

# Check status
sudo systemctl status devgraph
```

### Step 8: Configure Nginx

```bash
# Create Nginx configuration
sudo nano /etc/nginx/sites-available/devgraph
```

**Nginx configuration:**
```nginx
server {
    listen 80;
    server_name your-domain.com www.your-domain.com;

    # Frontend (React build)
    location / {
        root /opt/devgraph/frontend/dist;
        try_files $uri $uri/ /index.html;
        
        # Cache static assets
        location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg|woff|woff2|ttf)$ {
            expires 1y;
            add_header Cache-Control "public, immutable";
        }
    }

    # Backend API
    location /api/ {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Auth endpoints
    location /auth/ {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Increase max upload size for code submissions
    client_max_body_size 10M;
}
```

```bash
# Enable site and restart Nginx
sudo ln -s /etc/nginx/sites-available/devgraph /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx
```

### Step 9: Setup SSL with Let's Encrypt

```bash
# Obtain SSL certificate
sudo certbot --nginx -d your-domain.com -d www.your-domain.com

# Auto-renewal is configured by default
# Test renewal
sudo certbot renew --dry-run
```

### Step 10: Configure Frontend API URL

```bash
# Update frontend to use production API
nano /opt/devgraph/frontend/src/services/api.js
```

Change the baseURL to use relative paths (works with Nginx proxy):
```javascript
const API_BASE_URL = ''; // Empty for same origin (Nginx proxy)
```

Rebuild frontend:
```bash
cd /opt/devgraph/frontend
npm run build
```

### Step 11: Setup Firewall

```bash
# Configure UFW firewall
sudo ufw allow OpenSSH
sudo ufw allow 'Nginx Full'
sudo ufw enable

# Check status
sudo ufw status
```

### Step 12: Setup Monitoring & Logs

```bash
# View backend logs
sudo journalctl -u devgraph -f

# View Nginx logs
sudo tail -f /var/log/nginx/access.log
sudo tail -f /var/log/nginx/error.log
```

---

## 🐳 Docker Deployment

### Step 1: Create Dockerfile for Backend

Create `Dockerfile` in project root:

```dockerfile
# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build application
RUN CGO_ENABLED=0 GOOS=linux go build -o devgraph-server cmd/server/main.go

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/devgraph-server .
COPY --from=builder /app/.env.example .env

EXPOSE 8080

CMD ["./devgraph-server"]
```

### Step 2: Create Dockerfile for Frontend

Create `frontend/Dockerfile`:

```dockerfile
# Build stage
FROM node:18-alpine AS builder

WORKDIR /app

# Copy package files
COPY package*.json ./
RUN npm ci

# Copy source
COPY . .

# Build app
RUN npm run build

# Runtime stage
FROM nginx:alpine

# Copy built assets
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
```

### Step 3: Create Docker Compose File

Create `docker-compose.yml`:

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: devgraph
      POSTGRES_USER: devgraph_user
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U devgraph_user"]
      interval: 10s
      timeout: 5s
      retries: 5

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 10s
      timeout: 5s
      retries: 5

  backend:
    build:
      context: .
      dockerfile: Dockerfile
    environment:
      DB_HOST: postgres
      DB_PORT: 5432
      DB_USER: devgraph_user
      DB_PASSWORD: ${DB_PASSWORD}
      DB_NAME: devgraph
      JWT_SECRET: ${JWT_SECRET}
      SERVER_PORT: 8080
      REDIS_HOST: redis
      REDIS_PORT: 6379
      GIN_MODE: release
    ports:
      - "8080:8080"
    depends_on:
      postgres:
        condition: service_healthy
      redis:
        condition: service_healthy
    restart: unless-stopped

  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    ports:
      - "80:80"
    depends_on:
      - backend
    restart: unless-stopped

volumes:
  postgres_data:
  redis_data:
```

### Step 4: Create Frontend Nginx Config

Create `frontend/nginx.conf`:

```nginx
server {
    listen 80;
    server_name localhost;
    root /usr/share/nginx/html;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }

    location /api {
        proxy_pass http://backend:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /auth {
        proxy_pass http://backend:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

### Step 5: Create Docker Environment File

Create `.env.docker`:

```env
DB_PASSWORD=your_secure_db_password
JWT_SECRET=your_very_long_random_jwt_secret
```

### Step 6: Deploy with Docker Compose

```bash
# Build and start all services
docker-compose --env-file .env.docker up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Stop and remove volumes (WARNING: deletes data)
docker-compose down -v
```

### Step 7: Docker Production Deployment

For production with Docker on a VPS:

```bash
# Install Docker and Docker Compose on your server
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER

# Install Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Clone repository and deploy
git clone <your-repo-url> /opt/devgraph
cd /opt/devgraph
nano .env.docker  # Configure environment
docker-compose --env-file .env.docker up -d
```

---

## 🎯 Platform-Specific Guides

### AWS EC2 Deployment

1. **Launch EC2 Instance:**
   - Ubuntu 22.04 LTS
   - t2.medium or larger (2GB+ RAM)
   - Security Group: Allow ports 22, 80, 443

2. **Elastic IP:** Attach for static IP

3. **Follow "Production Deployment" steps above**

4. **RDS Setup (Optional):**
   - Use AWS RDS for PostgreSQL instead of local
   - Use ElastiCache for Redis

### DigitalOcean Droplet

1. **Create Droplet:**
   - Ubuntu 22.04
   - $12/month or higher
   - Add SSH key

2. **Follow "Production Deployment" steps**

3. **Use Managed Databases (Optional):**
   - Create PostgreSQL database
   - Use connection string in `.env`

### Heroku Deployment

1. **Install Heroku CLI**

2. **Create Heroku Apps:**
```bash
heroku create devgraph-backend
heroku create devgraph-frontend
```

3. **Add PostgreSQL and Redis:**
```bash
heroku addons:create heroku-postgresql:mini -a devgraph-backend
heroku addons:create heroku-redis:mini -a devgraph-backend
```

4. **Deploy Backend:**
```bash
# Create Procfile in root
echo "web: ./devgraph-server" > Procfile

# Set config vars
heroku config:set JWT_SECRET=your_secret -a devgraph-backend

# Deploy
git push heroku main
```

5. **Deploy Frontend:**
```bash
cd frontend
# Update api.js to use Heroku backend URL
heroku buildpacks:set heroku/nodejs -a devgraph-frontend
git subtree push --prefix frontend heroku main
```

### Vercel (Frontend) + Railway (Backend)

**Frontend on Vercel:**
1. Import repository to Vercel
2. Set root directory to `frontend`
3. Build command: `npm run build`
4. Output directory: `dist`

**Backend on Railway:**
1. Import repository to Railway
2. Add PostgreSQL and Redis services
3. Environment variables auto-configured
4. Update frontend API URL

---

## 🔒 Security Checklist

- [ ] Change all default passwords
- [ ] Generate strong JWT secret (32+ characters)
- [ ] Enable firewall (UFW/Security Groups)
- [ ] Setup SSL/TLS certificates
- [ ] Use environment variables (never commit .env)
- [ ] Regular security updates (`apt update && apt upgrade`)
- [ ] Setup database backups
- [ ] Configure rate limiting
- [ ] Use strong PostgreSQL password
- [ ] Disable unnecessary services
- [ ] Setup monitoring and alerts
- [ ] Configure CORS properly
- [ ] Use HTTPS only in production

---

## 🔄 Maintenance

### Updating Application

```bash
# Pull latest changes
cd /opt/devgraph
git pull

# Update backend
go build -o devgraph-server cmd/server/main.go
sudo systemctl restart devgraph

# Update frontend
cd frontend
npm install
npm run build
```

### Database Backup

```bash
# Backup PostgreSQL
pg_dump -U devgraph_user devgraph > backup_$(date +%Y%m%d).sql

# Restore
psql -U devgraph_user devgraph < backup_20260220.sql
```

### Monitoring

```bash
# Backend status
sudo systemctl status devgraph

# Resource usage
htop

# Disk space
df -h

# Backend logs
sudo journalctl -u devgraph -n 100 --no-pager
```

---

## ❓ Troubleshooting

### Backend won't start
- Check `.env` configuration
- Verify PostgreSQL is running: `systemctl status postgresql`
- Verify Redis is running: `systemctl status redis`
- Check logs: `journalctl -u devgraph -n 50`

### Database connection errors
- Verify credentials in `.env`
- Check PostgreSQL port: `sudo netstat -tlnp | grep 5432`
- Test connection: `psql -h localhost -U devgraph_user -d devgraph`

### Frontend can't connect to backend
- Check CORS settings in `cmd/server/main.go`
- Verify backend is running on port 8080
- Check `frontend/src/services/api.js` API URL
- Check browser console for errors

### Port already in use
```bash
# Find process using port 8080
sudo lsof -i :8080

# Kill process
sudo kill -9 <PID>
```

---

## 📞 Support

For issues and questions:
- Check existing documentation
- Review server logs
- Verify all services are running
- Check firewall rules

---

## 🎉 Success!

Your DevGraph application should now be deployed and accessible. Test all features:
- ✅ User registration
- ✅ Login/logout
- ✅ Code submission
- ✅ Code analysis
- ✅ Developer recommendations

Happy deploying! 🚀
