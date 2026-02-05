# DevGraph Frontend

Modern React frontend for the DevGraph code analysis platform.

## Features

✨ **Clean, Modern UI** - Built with React, Tailwind CSS, and Lucide icons
🔐 **Complete Authentication** - Register, login, logout with JWT tokens
💻 **Code Editor** - Monaco Editor with syntax highlighting for multiple languages
📊 **Code Analysis** - Submit code for automatic complexity and pattern analysis
👥 **Developer Recommendations** - Discover similar developers based on coding patterns
🎨 **Responsive Design** - Works beautifully on desktop and mobile

## Tech Stack

- **React 18** - Modern React with hooks
- **Vite** - Lightning-fast development server
- **Tailwind CSS** - Utility-first styling
- **Monaco Editor** - VS Code's editor for the web
- **Axios** - HTTP client with interceptors
- **React Router** - Client-side routing
- **React Hot Toast** - Beautiful notifications
- **Lucide React** - Beautiful icons

## Getting Started

### Prerequisites

- Node.js 18+ and npm

### Installation

```bash
cd frontend
npm install
```

### Development

```bash
npm run dev
```

The app will open at `http://localhost:3000`

### Build for Production

```bash
npm run build
npm run preview
```

## Project Structure

```
frontend/
├── src/
│   ├── components/      # Reusable components
│   │   ├── Navbar.jsx
│   │   └── ProtectedRoute.jsx
│   ├── context/         # React context
│   │   └── AuthContext.jsx
│   ├── pages/           # Page components
│   │   ├── Login.jsx
│   │   ├── Register.jsx
│   │   ├── Dashboard.jsx
│   │   └── Recommendations.jsx
│   ├── services/        # API services
│   │   └── api.js
│   ├── App.jsx          # Main app component
│   ├── main.jsx         # Entry point
│   └── index.css        # Global styles
├── index.html
├── package.json
├── vite.config.js
└── tailwind.config.js
```

## API Integration

The frontend connects to the Go backend at `http://localhost:8080`:

- `POST /auth/register` - User registration
- `POST /auth/login` - User login
- `POST /auth/refresh` - Token refresh
- `POST /auth/logout` - User logout
- `GET /api/me` - Get current user
- `POST /api/submit` - Submit code for analysis
- `GET /api/recommendations` - Get developer recommendations

## Environment Variables

The API base URL is configured in `src/services/api.js`. For production, update it to your backend URL.

## Features Overview

### Authentication
- Secure JWT-based authentication
- Automatic token refresh
- Protected routes

### Code Editor
- Monaco Editor (VS Code engine)
- Support for Python, JavaScript, Java, C++, Go
- Syntax highlighting and IntelliSense
- Submit code for analysis

### Developer Recommendations
- View developers with similar coding patterns
- Similarity scores and shared patterns
- Beautiful card-based UI

## Development Tips

1. **Backend CORS**: Ensure your Go backend allows CORS from `http://localhost:3000`
2. **Token Storage**: Tokens are stored in localStorage
3. **Auto-refresh**: Expired tokens are automatically refreshed
4. **Error Handling**: All API errors show user-friendly toast notifications

## Troubleshooting

**Problem**: API calls fail with CORS error
**Solution**: Add CORS middleware to your Go backend:
```go
import "github.com/gin-contrib/cors"

r.Use(cors.Default())
```

**Problem**: Monaco Editor not loading
**Solution**: Clear browser cache and restart dev server

## Contributing

This is the frontend for DevGraph. Make sure the backend is running on port 8080.

## License

MIT
