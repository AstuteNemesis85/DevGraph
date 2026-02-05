# DevGraph API Documentation

## Base URL
```
http://localhost:8080
```

---

## Authentication Endpoints

### Register
```http
POST /auth/register
Content-Type: application/json

{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (201 Created)**
```json
{
  "message": "user registered successfully",
  "user_id": "uuid"
}
```

---

### Login
```http
POST /auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (200 OK)**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "random_token_string"
}
```

---

### Refresh Token
```http
POST /auth/refresh
Content-Type: application/json

{
  "refresh_token": "your_refresh_token"
}
```

**Response (200 OK)**
```json
{
  "access_token": "new_access_token",
  "refresh_token": "new_refresh_token"
}
```

---

### Logout
```http
POST /auth/logout
Authorization: Bearer <access_token>
```

**Response (200 OK)**
```json
{
  "message": "logged out successfully"
}
```

---

## Protected Endpoints

All protected endpoints require the `Authorization` header:
```
Authorization: Bearer <access_token>
```

---

### Get Current User
```http
GET /api/me
Authorization: Bearer <access_token>
```

**Response (200 OK)**
```json
{
  "user_id": "uuid",
  "message": "you are authenticated"
}
```

---

### Submit Code
```http
POST /api/submit
Authorization: Bearer <access_token>
Content-Type: application/json

{
  "language": "python",
  "source_code": "def hello():\n    print('Hello, World!')"
}
```

**Supported Languages:**
- `python`
- `javascript`
- `java`
- `cpp`
- `go`

**Response (201 Created)**
```json
{
  "submission_id": "uuid",
  "message": "code submitted successfully"
}
```

---

### Get Recommendations
```http
GET /api/recommendations
Authorization: Bearer <access_token>
```

**Response (200 OK)**
```json
[
  {
    "id": "uuid",
    "user_a": "uuid",
    "user_b": "uuid",
    "similarity": 0.85,
    "shared_patterns": 5,
    "last_updated": "2025-12-24T10:30:00Z"
  }
]
```

**Fields:**
- `user_a`, `user_b`: UUIDs of the two similar users
- `similarity`: Similarity score (0.0 to 1.0)
- `shared_patterns`: Number of shared algorithm patterns
- `last_updated`: Last calculation timestamp

---

## Error Responses

### 400 Bad Request
```json
{
  "error": "validation error message"
}
```

### 401 Unauthorized
```json
{
  "error": "invalid credentials"
}
```

or

```json
{
  "error": "invalid refresh token"
}
```

### 409 Conflict
```json
{
  "error": "user already exists"
}
```

### 500 Internal Server Error
```json
{
  "error": "failed to store session"
}
```

---

## Token Lifecycle

1. **Login** → Receive `access_token` (15 min) + `refresh_token` (7 days)
2. **Use Access Token** → Include in `Authorization` header
3. **Token Expires** → Frontend auto-refreshes using `refresh_token`
4. **Refresh Success** → New tokens issued
5. **Refresh Fails** → User redirected to login

---

## Code Analysis Flow

1. User submits code via `POST /api/submit`
2. Code stored in database
3. Submission ID added to analysis queue
4. Background worker processes:
   - Detects algorithm patterns
   - Calculates time/space complexity
   - Identifies code issues
5. Updates graph with similarity calculations
6. Results available via `GET /api/recommendations`

---

## Rate Limiting

Currently not implemented. Recommended for production:
- 100 requests/minute per user
- 10 code submissions/hour per user

---

## CORS Configuration

Allowed origins:
- `http://localhost:3000` (development)

Allowed methods:
- GET, POST, PUT, PATCH, DELETE, OPTIONS

Allowed headers:
- Origin, Content-Type, Authorization

---

## Database Models

### User
```
id           UUID (PK)
username     String (unique)
email        String (unique)
password_hash String
created_at   Timestamp
```

### Session
```
id                UUID (PK)
user_id          UUID (FK)
refresh_token_hash String (indexed)
expires_at       Timestamp
created_at       Timestamp
```

### CodeSubmission
```
id          UUID (PK)
user_id     UUID (FK)
language    String
source_code Text
created_at  Timestamp
```

### CodeAnalysis
```
id               UUID (PK)
submission_id    UUID (FK, indexed)
time_complexity  String
space_complexity String
issues           String
created_at       Timestamp
```

### UserSimilarityEdge
```
id             UUID (PK)
user_a         UUID (indexed)
user_b         UUID (indexed)
similarity     Float (0.0-1.0)
shared_patterns Int
last_updated   Timestamp
```

---

## Redis Cache

### Session Cache
```
Key: "session:<refresh_token_hash>"
Value: "<user_id>"
TTL: 7 days
```

---

## Testing with cURL

### Register
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","email":"test@example.com","password":"password123"}'
```

### Login
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

### Submit Code
```bash
curl -X POST http://localhost:8080/api/submit \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -d '{"language":"python","source_code":"print(\"Hello\")"}'
```

### Get Recommendations
```bash
curl http://localhost:8080/api/recommendations \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

---

## WebSocket Support

Not currently implemented. Future consideration for real-time features:
- Live code analysis updates
- Real-time collaboration
- Notification system
