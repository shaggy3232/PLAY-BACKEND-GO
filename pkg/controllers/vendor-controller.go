package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/models"
	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/utils"
)

var NewVendor models.Vendor

func GetVendor(w http.ResponseWriter, r *http.Request) {
	newVendor := models.GetAllVendors()
	res, _ := json.Marshal(newVendor)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetVendorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendorId := vars["vendorId"]
	ID, err := strconv.ParseInt(vendorId, 0, 0)
	if err != nil {
		fmt.Println("ERROR WHILE PARSING ")
	}
	vendorDetails, _ := models.GetVendorById(ID)
	res, _ := json.Marshal(vendorDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateVendor(w http.ResponseWriter, r *http.Request) {
	CreateVendor := &models.Vendor{}
	utils.ParseBody(r, CreateVendor)
	v := CreateVendor.CreateVendor()
	res, _ := json.Marshal(v)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteVendor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendorId := vars["vendorId"]
	ID, err := strconv.ParseInt(vendorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	vendor := models.DeleteVendorById(ID)
	res, _ := json.Marshal(vendor)
	w.Header().Set("Content-Type", "pkglication.json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateVendor(w http.ResponseWriter, r *http.Request) {
	var updateVendor = &models.Vendor{}
	utils.ParseBody(r, updateVendor)
	vars := mux.Vars(r)
	vendorId := vars["vendorId"]
	ID, err := strconv.ParseInt(vendorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	vendorDetails, db := models.GetVendorById(ID)

	if updateVendor.Name != "" {
		vendorDetails.Name = updateVendor.Name
	}

	if updateVendor.Price != vendorDetails.Price {
		vendorDetails.Price = updateVendor.Price
	}

	if updateVendor.TravelingDistance != vendorDetails.TravelingDistance {
		vendorDetails.TravelingDistance = updateVendor.TravelingDistance
	}

	db.Save(&vendorDetails)
	res, _ := json.Marshal(vendorDetails)
	w.Header().Set("Content-Typer", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
