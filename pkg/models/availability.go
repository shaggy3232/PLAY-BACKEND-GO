package models

import (
	"time"

	"gorm.io/gorm"
)

type Availability struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	VendorID uint
	Entries  []AvailabilityEntry `json:"entries" gorm:"foreignKey:AvailabilityID"`
}

type AvailabilityEntry struct {
	ID             uint `gorm:"primaryKey"`
	AvailabilityID uint
	Day            time.Weekday `json:"day"`
	Start          int          `json:"start"`
	End            int          `json:"end"`
}
