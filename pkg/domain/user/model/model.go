package model

import (
	"gorm.io/gorm"
)

// User is the model for the DB
type User struct {
	gorm.Model
	Email    string `gorm:"unique_index;not null" json:"email"`
	Name     string `json:"name"`
	Password string `gorm:"not null" json:"password,omitempty"`
}

// AuthRequest handles data for /auth/login
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse is the response DTO for the auth resource
type AuthResponse struct {
	Token string `json:"token"`
}

// UserResponse handles data for POST in /user
type Response struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// UserRequest handles data for PUT/POST in /users
type Request struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
