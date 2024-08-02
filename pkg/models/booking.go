package models

type Booking struct {
	ID       uint   `db:"id"`
	VendorID uint   `db:"vendor_id"`
	User     string `db:"user"`
	Date     string `db:"date"`
	Start    int    `db:"start"`
	End      int    `db:"end"`
	Price    int    `db:"price"`
	Location string `db:"location"`
	Accepted bool   `db:"accepted"`
}
