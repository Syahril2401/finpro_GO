# API Curl Examples

### 1. Authentication

**Register**
```bash
curl -X POST http://localhost:8080/api/auth/register \
-H "Content-Type: application/json" \
-d '{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}'
```

**Login**
```bash
curl -X POST http://localhost:8080/api/auth/login \
-H "Content-Type: application/json" \
-d '{
  "email": "john@example.com",
  "password": "password123"
}'
```

### 2. Assessment (Requires Token)

**Get Questions**
```bash
curl -X GET http://localhost:8080/api/assessment/questions \
-H "Authorization: Bearer <YOUR_TOKEN>"
```

**Submit Assessment**
```bash
curl -X POST http://localhost:8080/api/assessment/submit \
-H "Authorization: Bearer <YOUR_TOKEN>" \
-H "Content-Type: application/json" \
-d '[
  {"question_id": 1, "answer": "Sangat Sering", "score": 5},
  {"question_id": 2, "answer": "Sering", "score": 4}
]'
```

### 3. Dashboard (Requires Token)

**Get Dashboard Data**
```bash
curl -X GET http://localhost:8080/api/dashboard \
-H "Authorization: Bearer <YOUR_TOKEN>"
```

### 4. Admin (Requires Admin Token)

**Get Users**
```bash
curl -X GET http://localhost:8080/api/admin/users \
-H "Authorization: Bearer <ADMIN_TOKEN>"
```

**Get Evaluations**
```bash
curl -X GET http://localhost:8080/api/admin/evaluations \
-H "Authorization: Bearer <ADMIN_TOKEN>"
```
