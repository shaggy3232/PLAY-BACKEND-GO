package models

import "time"

type Booking struct {
	ID          string    `json:"id"` // string representation of UUID
	RefereeID   string    `json:"referee_id"`
	OrganizerID string    `json:"organizer_id"`
	Price       float32   `json:"price"`
	Start       time.Time `json:"start_time"`
	End         time.Time `json:"end_time"`
	Location    string    `json:"location"`
	Accepted    bool      `json:"accepted"`
	Cancelled   bool      `json:"canceled"`
	LastUpdated time.Time `json:"last_updated"`
	CreatedAt   time.Time `json:"created_at"`
}
