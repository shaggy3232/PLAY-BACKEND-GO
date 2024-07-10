package models

import (
	"time"

	"gorm.io/gorm"
)

type Availability struct {
	gorm.Model
	Entries []AvailabilityEntry `json:"entries"`
}

type AvailabilityEntry struct {
	Day   time.Weekday `json:"day"`
	Start int          `json:"start"`
	End   int          `json:"end"`
}
