package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Availability struct {
	ID       uint                `db:"id"`
	VendorID uint                `json:"vendor_id"`
	Entries  []AvailabilityEntry `json:"entries"`
}

type AvailabilityEntry struct {
	ID             uint   `json:"id" db:"id"`
	AvailabilityID uint   `json:"availability_id"`
	Day            string `json:"day"`
	Start          int    `json:"start"`
	End            int    `json:"end"`
}
