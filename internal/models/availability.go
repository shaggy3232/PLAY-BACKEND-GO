package models

type Availability struct {
	///db schema
	ID     string  `json:"id"` // string representation of UUID
	UserID string  `json:"user_id"`
	Price  float32 `json:"price"`
	Start  string  `json:"start"`
	End    string  `json:"end"`
}
