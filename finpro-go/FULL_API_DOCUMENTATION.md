# SRL Platform Backend REST API

Backend API for Self Regulated Learning Platform built with Go (Gin), GORM (MySQL), and Clean Architecture.

## 🚀 How to Run

1.  **Requirement**: Go 1.20+, MySQL.
2.  **Database Setup**:
    *   Create a database named `srl_platform` in your MySQL.
    *   Import the schema from `migrations/pt_database.sql`.
3.  **Environment**: 
    *   Adjust `.env` file with your MySQL credentials.
4.  **Run Application**:
    ```bash
    go run cmd/server/main.go
    ```
    The server will start on `http://localhost:8080`.

## 📂 Project Structure

*   `cmd/server/`: Main application entry point.
*   `config/`: Database and environment configuration.
*   `model/`: Database models and request/response structures.
*   `repository/`: Data access layer (Interfaces & Implementations).
*   `service/`: Business logic layer.
*   `controller/`: HTTP handlers.
*   `middleware/`: Auth and CORS middleware.
*   `routes/`: API route definitions.
*   `utils/`: Helper functions (JWT, Hash, Response).

## 🔒 Security

*   Passwords are hashed using **Bcrypt**.
*   Route protection using **JWT** (expires in 24h).
*   **RBAC** (Role Based Access Control) for Admin routes.
*   **Database Transactions** for assessment submissions.

# REST API Reference - Lumora App


## 🔒 Autentikasi (Authorization)
Sebagian besar endpoint memerlukan token agar dapat diakses. Authentication yang digunakan adalah **Bearer Token JWT**.
- Key: `Authorization`
- Value: `Bearer <token_dari_proses_login>`

---

## 1. Authentication (`/api/auth`)

### 1.1 Register User
Endpoint untuk mendaftarkan akun pengguna baru.
- **Method:** `POST`
- **Route:** `/api/auth/register`
- **Auth Required:** ❌
- **Request Body (JSON):**
  ```json
  {
    "name": "Alex J.",
    "email": "alex@example.com",
    "password": "mypassword123"
  }
  ```
- **Response (201 Created):**
  ```json
  {
    "success": true,
    "message": "Registration successful",
    "data": {
      "token": "eyJhb..."
    }
  }
  ```

### 1.2 Login User
Endpoint untuk masuk ke dalam sistem dan mendapatkan Token.
- **Method:** `POST`
- **Route:** `/api/auth/login`
- **Auth Required:** ❌
- **Request Body (JSON):**
  ```json
  {
    "email": "alex@example.com",
    "password": "mypassword123"
  }
  ```
- **Response (200 OK):**
  ```json
  {
    "success": true,
    "message": "Login successful",
    "data": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5c..."
    }
  }
  ```

---

## 2. Admin System (`/api/admin`)

### 2.1 Admin Login
Digunakan khusus untuk admin agar mendapatkan admin token.
- **Method:** `POST`
- **Route:** `/api/admin/login`
- **Auth Required:** ❌
- **Request Body (JSON):**
  ```json
  {
      "email": "admin@example.com",
      "password": "admin123"
  }
  ```

### 2.2 Get Users List
Mengambil seluruh daftar user terdaftar (Khusus Admin).
- **Method:** `GET`
- **Route:** `/api/admin/users`
- **Auth Required:** ✅ (Admin Token req)

### 2.3 Get Evaluations List
Mengambil semua hasil evaluasi dan assessment user.
- **Method:** `GET`
- **Route:** `/api/admin/evaluations`
- **Auth Required:** ✅ (Admin Token req)

---

## 3. Assessment (`/api/assessment`)

### 3.1 Get Initial Questions
Mengambil seluruh daftar pertanyaan assessment.
- **Method:** `GET`
- **Route:** `/api/assessment/questions`
- **Auth Required:** ✅ (Bearer Token)

### 3.2 Submit Assessment
Mengirimkan keseluruhan jawaban assessment user.
- **Method:** `POST`
- **Route:** `/api/assessment/submit`
- **Auth Required:** ✅ (Bearer Token)
- **Request Body (JSON):**
  ```json
  [
    {
      "question_id": "f47ac10b-58cc-...",
      "answer_value": 4
    },
    {
      "question_id": "123dc10c-58cc-...",
      "answer_value": 2
    }
  ]
  ```

---

## 4. Dashboard & Features (`/api/dashboard`)
> Seluruh endpoint di grup ini wajib menggunakan **Bearer Token** pada Authorization Headernya.

### 4.1 Dashboard Overview
Menampilkan statistik umum / overview untuk widget dashboard.
- **Method:** `GET`
- **Route:** `/api/dashboard`
- **Auth Required:** ✅

### 4.2 Planner View (Mock)
Menampilkan data agenda rutin/planner.
- **Method:** `GET`
- **Route:** `/api/dashboard/planner`
- **Auth Required:** ✅

### 4.3 Notes (Mock)
Menampilkan daftar catatan dan file materi pengguna.
- **Method:** `GET`
- **Route:** `/api/dashboard/notes`
- **Auth Required:** ✅

### 4.4 Weekly Targets (Mock)
Mengambil data progres target belajar mingguan.
- **Method:** `GET`
- **Route:** `/api/dashboard/weekly-targets`
- **Auth Required:** ✅

### 4.5 AI Strategies (Mock)
Mengambil rekomendasi strategi belajar yang di-generate oleh AI (Contoh: Pomodoro technique).
- **Method:** `GET`
- **Route:** `/api/dashboard/ai-strategies`
- **Auth Required:** ✅

### 4.6 Progress (Mock)
Menampilkan progress tracker.
- **Method:** `GET`
- **Route:** `/api/dashboard/progress`
- **Auth Required:** ✅

### 4.7 Settings (Mock)
Mengatur konfigurasi profil / user account.
- **Method:** `GET`
- **Route:** `/api/dashboard/settings`
- **Auth Required:** ✅

---

## 5. Test Routes (`/api/test`)
Endpoint sementara yang digunakan untuk penulisan standar saat *Basic Task*.
- **`GET`** `/api/test/add` - Tambah data via URL Param
- **`POST`** `/api/test/add` - Tambah data via JSON payload
- **`GET`** `/api/test/list` - Menampilkan data test


