package models

import (
	"database/sql/driver"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

type Availability struct {
	ID       uint
	VendorID uint                `json:"vendor_id"`
	Entries  []AvailabilityEntry `json:"entries"`
}

type AvailabilityEntry struct {
	ID             uint
	AvailabilityID uint   `json:"availability_id"`
	Day            string `json:"day"`
	Start          string `json:"start"`
	End            string `json:"end"`
}

func (a *Availability) Value() (driver.Value, error) {
	if a != nil {
		b, err := json.Marshal(a)
		if err != nil {
			return nil, err
		}
		return string(b), nil
	}
	return nil, nil
}

func (a *Availability) Scan(src interface{}) error {
	var data []byte
	if b, ok := src.([]byte); ok {
		data = b
	} else if s, ok := src.(string); ok {
		data = []byte(s)
	}
	return json.Unmarshal(data, a)
}
