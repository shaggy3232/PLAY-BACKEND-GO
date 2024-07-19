package models

import (
	"database/sql"
	"fmt"

	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/config"
)

var db *sql.DB

type Vendor struct {
	ID                uint         `json:"id" db:"id"`
	Name              string       `json:"name"`
	Price             int          `json:"price"`
	TravelingDistance int          `json:"travelingdistance"`
	Availability      Availability `json:"availability" gorm:"foreignKey:VendorID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	// BookingRequest    BookingRequest `json:"bookingrequest"`
	// Bookings          []Booking      `json:"bookings"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	createTables(db)

}

// Create Tables function
func createTables(db *sql.DB) error {
	vendorsTable := `CREATE TABLE IF NOT EXISTS vendors (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price INT,
		traveling_distance INT
	);`

	availabilitiesTable := `CREATE TABLE IF NOT EXISTS availabilities (
		id INT AUTO_INCREMENT PRIMARY KEY,
		vendor_id INT NOT NULL,
		FOREIGN KEY (vendor_id) REFERENCES vendors(id)
	);`

	availabilityEntriesTable := `CREATE TABLE IF NOT EXISTS availability_entries (
		id INT AUTO_INCREMENT PRIMARY KEY,
		availability_id INT NOT NULL,
		day VARCHAR(10) NOT NULL,
		start_time VARCHAR(8) NOT NULL,
		end_time VARCHAR(8) NOT NULL,
		FOREIGN KEY (availability_id) REFERENCES availabilities(id)
	);`

	_, err := db.Exec(vendorsTable)
	if err != nil {
		return err
	}
	_, err = db.Exec(availabilitiesTable)
	if err != nil {
		return err
	}
	_, err = db.Exec(availabilityEntriesTable)
	return err
}
func (v *Vendor) CreateVendor() *Vendor {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("ERROR WHILE Starting transaction", err)
	}

	result, err := tx.Exec("INSERT INTO vendors (name,price,traveling_distance) VALUES (?,?,?)", v.Name, v.Price, v.TravelingDistance)
	if err != nil {
		tx.Rollback()
		fmt.Println("ERROR WHILE Inserting name and price to vendor", err)
	}
	vendorID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println("ERROR WHILE Inserting to vendors")
	}

	v.ID = uint(vendorID)

	result, err = tx.Exec("INSERT INTO availabilities (vendor_id) VALUES (?)", v.ID)
	if err != nil {
		tx.Rollback()
		fmt.Println("ERROR WHILE Inserting to availabilities")
	}

	availabilityID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println("ERROR WHILE Inserting to availabilities")
	}
	v.Availability.ID = uint(availabilityID)
	v.Availability.VendorID = v.ID
	for _, entry := range v.Availability.Entries {
		_, err = tx.Exec("INSERT INTO availability_entries (availability_id, day, start_time, end_time) VALUES (?, ?, ?, ?)",
			uint(availabilityID), entry.Day, entry.Start, entry.End)
		fmt.Println(v.Availability.ID, availabilityID)
		if err != nil {
			tx.Rollback()
			fmt.Println("ERROR WHILE Inserting to availabilities entries")
		}
	}
	return v
}

func GetAllVendors() []Vendor {
	var vendors []Vendor

	return vendors
}

func GetVendorById(Id int64) (*Vendor, *sql.DB) {
	var getVendor Vendor

	return &getVendor, db
}

func DeleteVendorById(ID int64) Vendor {
	var vendor Vendor

	return vendor

}

func GetAllAvailibleVendors(day string, start string, end string) []Vendor {
	var availbleVendors []Vendor
	// Join the tables and apply the conditions

	return availbleVendors
}

func VendorIsAvailble(start string, end string, day string, vendor Vendor) bool {

	for _, entry := range vendor.Availability.Entries {
		if entry.Day == day {
			if entry.Start <= start && entry.End >= end {
				return true
			}
		}
	}
	return false
}
