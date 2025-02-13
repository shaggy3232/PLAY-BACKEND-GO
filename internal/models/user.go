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
