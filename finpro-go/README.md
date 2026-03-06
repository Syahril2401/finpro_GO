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
