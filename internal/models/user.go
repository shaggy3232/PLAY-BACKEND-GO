package models

import "time"

type User struct {
	///db schema
	ID          string    `json:"id"` // string representation of UUID
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	UserRole    string    `json:"user_role"`
	CreatedAt   time.Time `json:"created_at"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	token   string
	message string
}
