package models

import (
	"time"

	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Vendor struct {
	gorm.Model
	ID                uint           `gorm:"primaryKey"`
	Name              string         `json:"name"`
	Price             int            `json:"price"`
	TravelingDistance int            `json:"travelingdistance"`
	Availability      Availability   `json:"availability" gorm:"foreignKey:VendorID"`
	BookingRequest    BookingRequest `json:"bookingrequest"`
	Bookings          []Booking      `json:"bookings"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Vendor{}, &Availability{}, &AvailabilityEntry{})
}

func (v *Vendor) CreateVendor() *Vendor {
	db.Create(&v)
	return v
}

func GetAllVendors() []Vendor {
	var Vendors []Vendor
	db.Find(&Vendors)
	return Vendors
}

func GetVendorById(Id int64) (*Vendor, *gorm.DB) {
	var getVendor Vendor
	db := db.Where("ID=?", Id).Find(&getVendor)
	return &getVendor, db
}

func DeleteVendorById(Id int64) Vendor {
	var vendor Vendor
	db.Where("ID = ?", Id).Delete(vendor)
	return vendor
}

func GetAllAvailibleVendors(start time.Time, end time.Time) []Vendor {
	var availbleVendors []Vendor
	// Join the tables and apply the conditions
	err := db.Joins("JOIN availabilities ON availabilities.vendor_id = vendors.id").
		Joins("JOIN availability_entries ON availability_entries.availability_id = availabilities.id").
		Where("availability_entries.day = ? AND availability_entries.start_time <= ? AND availability_entries.end_time >= ?", start.Weekday(), start.Hour(), end.Hour()).
		Preload("Availability.AvailabilityEntries").
		Find(&availbleVendors).Error
	if err != nil {
		return nil
	}

	return availbleVendors
}

func VendorIsAvailble(start time.Time, end time.Time, vendor Vendor) bool {
	day := start.Weekday()

	for _, entry := range vendor.Availability.Entries {
		if entry.Day == day {
			if entry.Start <= start.Hour() && entry.End >= end.Hour() {
				return true
			}
		}
	}
	return false
}
