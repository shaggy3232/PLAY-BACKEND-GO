package models

import (
	"time"
)

type Availability struct {
	///db schema
	ID        string    `json:"id"` // string representation of UUID
	UserID    string    `json:"user_id"`
	Price     float32   `json:"price"`
	Start     time.Time `json:"start_time"`
	End       time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
}
