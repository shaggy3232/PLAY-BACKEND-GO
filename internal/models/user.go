package models

type User struct {
	///db schema
	ID          string `json:"id"` // string representation of UUID
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type UserList struct {
	Users []*User
}
