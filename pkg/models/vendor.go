package models

import (
	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Vendor struct {
	gorm.Model
	Name              string `gorm:"" json:"name"`
	Price             int    `json:"price"`
	TravelingDistance int    `json:"travelingdistance"`
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
