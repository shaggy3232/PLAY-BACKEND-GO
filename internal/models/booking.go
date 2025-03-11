package models

type Booking struct {
	ID          string  `json:"id"` // string representation of UUID
	RefereeID   string  `json:"referee_id"`
	OrganizerID string  `json:"organizer_id"`
	Price       float32 `json:"price"`
	Start       string  `json:"start"`
	End         string  `json:"end"`
	Location    string  `json:"location"`
	Accepted    bool    `json:"accepted"`
	Cancelled   bool    `json:"canceled"`
	LastUpdated string  `json:"last_updated"`
	CreatedAt   string  `json:"created_at"`
}
