package models

type Availbility struct {
	///db schema
	ID     string  `json:"id"` // string representation of UUID
	UserID string  `json:"user_id"`
	Price  float32 `json:"price"`
	Start  string  `json:"start"`
	End    string  `json:"end"`
}

type AvailbilityList struct {
	Availbilities []*Availbility
}
