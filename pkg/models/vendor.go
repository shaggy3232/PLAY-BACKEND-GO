package models

import (
	"time"

	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Vendor struct {
	gorm.Model
	Name              string         `gorm:"" json:"name"`
	Price             int            `json:"price"`
	TravelingDistance int            `json:"travelingdistance"`
	Availability      Availability   `json:"availability"`
	BookingRequest    BookingRequest `json:"bookingrequest"`
	Bookings          []Booking      `json:"bookings"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Vendor{})
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
	err := db.Preload("Availability", "day = ? AND start_time <= ? AND end_time >= ?", start.Weekday(), start.Hour(), end.Hour()).
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
