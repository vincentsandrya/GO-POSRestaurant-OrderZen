package dto

import "time"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseLogin struct {
	UserId      uint      `json:"user_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	RoleId      uint      `json:"role_id"`
	CreatedDate time.Time `json:"created_date"`
	Token       string    `json:"token"`
}

type UserRequest struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	RoleId      uint      `json:"role_id"`
	CreatedDate time.Time `json:"created_date"`
	CreatedBy   string    `json:"created_by"`
	UpdatedDate time.Time `json:"updated_date"`
	UpdatedBy   string    `json:"updated_by"`
}

type UserResponse struct {
	UserId      uint      `json:"user_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	RoleId      uint      `json:"role_id"`
	CreatedDate time.Time `json:"created_date"`
	CreatedBy   string    `json:"created_by"`
	UpdatedDate time.Time `json:"updated_date"`
	UpdatedBy   string    `json:"updated_by"`
}
