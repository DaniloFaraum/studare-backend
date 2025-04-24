# 🎓 Studare – Course Evaluation Platform

This repository contains the source code and documentation for Studare, a course evaluation platform. The platform aims to provide a space where students can evaluate courses in a structured and meaningful way, offering valuable feedback to improve academic quality.

---

  ## ✨ Features:

The Studare platform includes the following core features:

  ### 🔐 Authentication & User Management
  - User Login – Secure login system with credential verification.

  - JWT Authentication – Uses JSON Web Tokens to manage user sessions and protect private routes.

  - Role-Based Access Control – Different access levels for admins and regular users.

  ### 🎓 Course Management
  - Create Courses – Admins can add new courses with relevant details.
  
  - Read/View Courses – All users can browse and view course information.
  
  - Update Courses – Admins can edit course data at any time.
  
  - Delete Courses – Admins can remove outdated or invalid courses.

  ### 📝 Rating System
  - Submit Ratings – Authenticated users can submit feedback and rate courses.
  
  - Edit Ratings – Users can update their own evaluations.
  
  - Delete Ratings – Users can delete their feedback if necessary.
  
  - View Aggregated Ratings – See overall course ratings based on user input.

  - Custom Rating Algorithm - Algorithm with more precision of real public opinion 

---

## 🚀 How to Run

1. Clone this repository:
   ```bash
   git clone https://github.com/DaniloFaraum/studare-backend
   cd Java-Solid
2. Import it into your favorite IDE (Eclipse, IntelliJ, or VS Code with Java support).
3. Download the dependencies:
    ```
    go mod tidy
   ```
4. Start the server:
    ```
    cd cmd
    go run main.go
    ```

## 🛠 Technologies Used

Backend: Go, Gin Gonic, GORM 

Database: MySQL, Azure

Tools: Visual Studio Code, Git & GitHub, Postman, Notion

## 📚 References
Go Language Documentation – https://golang.org/doc

Gin Web Framework – https://gin-gonic.com

Clean Architecture – Robert C. Martin
