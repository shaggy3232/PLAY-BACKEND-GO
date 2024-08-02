package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/config"
)

var db *sql.DB

type Vendor struct {
	ID                uint         `json:"id" db:"id"`
	Name              string       `json:"name"`
	Price             int          `json:"price"`
	TravelingDistance int          `json:"travelingdistance"`
	Availability      Availability `json:"availability"`
	Bookings          []Booking    `json:"bookings"`
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
			ON DELETE CASCADE
	);`

	availabilityEntriesTable := `CREATE TABLE IF NOT EXISTS availability_entries (
		id INT AUTO_INCREMENT PRIMARY KEY,
		availability_id INT NOT NULL,
		day VARCHAR(10) NOT NULL,
		start_time INT NOT NULL,
		end_time INT NOT NULL,
		FOREIGN KEY (availability_id) REFERENCES availabilities(id)
			ON DELETE CASCADE
	);`

	bookingsTable := `CREATE TABLE IF NOT EXISTS bookings (
		id INT AUTO_INCREMENT PRIMARY KEY,
		vendor_id INT NOT NULL,
		user VARCHAR(255) NOT NULL,
		date VARCHAR(10) NOT NULL,
		start INT NOT NULL,
		end INT NOT NULL,
		price int NOT NULL,
		location VARCHAR(255),
		is_accepted BOOL,
		FOREIGN KEY (vendor_id) REFERENCES vendors(id)
			ON DELETE CASCADE
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
	if err != nil {
		return err
	}

	_, err = db.Exec(bookingsTable)
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
	fmt.Println(v.Availability.ID, availabilityID)
	for i, entry := range v.Availability.Entries {
		result, err = tx.Exec("INSERT INTO availability_entries (availability_id, day, start_time, end_time) VALUES (?, ?, ?, ?)",
			v.Availability.ID, entry.Day, entry.Start, entry.End)
		v.Availability.Entries[i].AvailabilityID = v.Availability.ID
		availEntryId, _ := result.LastInsertId()
		v.Availability.Entries[i].ID = uint(availEntryId)
		if err != nil {
			tx.Rollback()
			fmt.Println("ERROR WHILE Inserting to availabilities entries")
		}
	}
	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		fmt.Print(err)
	}
	var bookings []Booking
	v.Bookings = bookings

	return v
}

func GetAllVendors() []Vendor {
	var vendors []Vendor
	availabilities, err := GetAllAvailabilities()
	if err != nil {
		fmt.Print(err)
	}

	query := "SELECT * from vendors"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Print(err)
	}

	for rows.Next() {
		var vendor Vendor
		if err := rows.Scan(&vendor.ID, &vendor.Name, &vendor.Price, &vendor.TravelingDistance); err != nil {
			fmt.Print(err)
			return nil
		}

		for _, a := range availabilities {
			if a.VendorID == vendor.ID {
				vendor.Availability = a
			}
		}
		vendors = append(vendors, vendor)
	}

	return vendors
}

func GetVendorById(Id int64) Vendor {
	query := "SELECT * FROM vendors WHERE id = ?"
	row := db.QueryRow(query, Id)
	var vendor Vendor

	row.Scan(&vendor.ID, &vendor.Name, &vendor.Price, &vendor.TravelingDistance)
	availability, err := GetAllAvailabilitiesWithVendorID(int(vendor.ID))
	if err != nil {
		fmt.Print(err)
	}

	vendor.Availability = availability
	return vendor

}

func DeleteVendorById(vendor Vendor) Vendor {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("ERROR WHILE Starting transaction", err)
	}
	_, er := tx.Exec("DELETE FROM vendors WHERE id = ?", vendor.ID)
	if er != nil {
		tx.Rollback()
		fmt.Println("ERROR WHILE Deleting name and price to vendor", er)
	}
	tx.Commit()

	return vendor

}

func GetAllAvailibleVendors(day string, start string, end string) []Vendor {
	var availbleVendors []Vendor
	// Join the tables and apply the conditions

	return availbleVendors
}

func VendorIsAvailble(start int, end int, dateString string, vendor Vendor) bool {
	inAvailability := false
	noBookings := true
	// date format is YYYY-MM-DD
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		fmt.Println("ERROR WHILE PARSING DATE", err)
	}
	day := date.Weekday().String()

	for _, entry := range vendor.Availability.Entries {
		if entry.Day == day {
			if entry.Start <= start && entry.End >= end {
				inAvailability = true
			}
		}
	}
	for _, booking := range vendor.Bookings {
		if booking.Date == dateString {
			if booking.Start <= end || booking.End >= start {
				noBookings = false
			}
		}
	}
	return noBookings && inAvailability
}
func requestBooking(start int, end int, dateString string, vendorID int, user string, location string, Price int) Booking {
	var booking Booking
	booking.VendorID = uint(vendorID)
	booking.User = user
	booking.Start = start
	booking.End = end
	booking.Date = dateString
	booking.Price = Price
	booking.Location = location
	booking.Accepted = false

	vendor := GetVendorById(int64(vendorID))

	isVendorAvailable := VendorIsAvailble(start, end, dateString, vendor)

	if isVendorAvailable {
		CreateBooking(booking)
	}
	return booking

}
func CreateBooking(booking Booking) {
	query := "INSERT INTO bookings (vendor_id, user, date, start, end, price, location, is_accepted) VALUES (?,?,?,?,?,?,?,?)"
	result, err := db.Exec(query, booking.VendorID, booking.User, booking.Date, booking.Start, booking.End, booking.Price, booking.Location, booking.Accepted)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(result)

}

func GetAllAvailabilityEntries() ([]AvailabilityEntry, error) {
	var availabilityEntries []AvailabilityEntry
	query := "SELECT * FROM availability_entries"

	rows, err := db.Query(query)

	for rows.Next() {
		var entry AvailabilityEntry
		if err := rows.Scan(&entry.ID, &entry.AvailabilityID, &entry.Day, &entry.Start, &entry.End); err != nil {
			return nil, err
		}
		availabilityEntries = append(availabilityEntries, entry)
	}

	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return availabilityEntries, nil
}

func GetAllAvailabilities() ([]Availability, error) {
	var availabilities []Availability
	availabilityentries, err := GetAllAvailabilityEntries()
	if err != nil {
		fmt.Print(err)
	}
	query := "SELECT * FROM availabilities"

	rows, err := db.Query(query)

	for rows.Next() {
		var availability Availability
		if err := rows.Scan(&availability.ID, &availability.VendorID); err != nil {
			return nil, err
		}
		for _, s := range availabilityentries {
			if s.AvailabilityID == availability.ID {
				availability.Entries = append(availability.Entries, s)
			}
		}
		availabilities = append(availabilities, availability)
	}

	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return availabilities, nil
}

func GetAllAvailabilitiesWithVendorID(ID int) (Availability, error) {
	query := "SELECT * FROM availabilities WHERE vendor_id = ?"
	row := db.QueryRow(query, ID)
	var availability Availability
	err := row.Scan(&availability.ID, &availability.VendorID)
	if err != nil {
		return availability, err
	}
	availabilities := GetAllEntriesGivenAvailabilityID(int(availability.ID))
	availability.Entries = availabilities
	return availability, err
}

func GetAllEntriesGivenAvailabilityID(ID int) []AvailabilityEntry {
	var availabilities []AvailabilityEntry
	query := "SELECT * FROM availability_entries WHERE availability_id = ?"
	rows, err := db.Query(query, ID)
	if err != nil {
		fmt.Print(err)
	}
	for rows.Next() {
		var entry AvailabilityEntry
		if err := rows.Scan(&entry.ID, &entry.AvailabilityID, &entry.Day, &entry.Start, &entry.End); err != nil {
			fmt.Print(err)
		}
		availabilities = append(availabilities, entry)
	}
	return availabilities
}
