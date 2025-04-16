# 🧾 OrderZen - POS System in Go

OrderZen is a modern, efficient, and lightweight Point of Sale (POS) system built with **Golang**. Designed for speed, reliability, and ease of use, it helps small and medium businesses manage sales, inventory, and customer transactions seamlessly.

---

## 🚀 Features

- 🧑‍💼 User authentication with JWT
- 🛍️ Menu and category management
- 🧾 Order and transaction processing
- 💰 Discount and promo code system
- 📊 Basic reporting (sales, etc.)
- 🛡️ Secure API endpoints
- 📁 Modular project structure

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin / Gorm
- **Database:** PostgreSQL
- **ORM:** GORM
- **Authentication:** JWT
- **Environment Config:** GoDotEnv

---

## 📁 Project Structure

```bash
GO-POSRestaurant-OrderZen/
├── database/         # Connection to Postgres
├── display/          # List of template response
├── dto/              # Data Transfer Objects
├── helpers/          # Function hashing and mapping dto become model
├── internal/         # for handling, service, and repository (data access logic)
├── middleware/       # JWT and other middleware
├── models/           # Database models
├── main.go           # Main entry file
├── go.mod            # Go modules
└── README.md         # Project documentation


## 🛠️ Documentation API
**auth
- POST /login
- POST /register

**user
- GET /user
- GET /user/:id
- POST /user
- PUT /user/:id
- DELETE /user/:id

**category menu
- GET /menu/category
- GET /menu/category/:id
- POST /menu/category
- PUT /menu/category/:id
- DELETE /menu/category/:id

**menu
- GET /menu
- GET /menu/:id
- POST /menu
- PUT /menu/:id
- DELETE /menu/:id

**promo
- GET /promo
- GET /promo/:id
- POST /promo
- PUT /promo/:id
- DELETE /promo/:id

**order
- GET /promo
- GET /promo/:id
- POST /promo

**payment
- GET /payment
- GET /payment/:id
