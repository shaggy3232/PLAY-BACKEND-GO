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
	vendorDetails := models.GetVendorById(ID)
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
	fmt.Println(ID)
	if err != nil {
		fmt.Println("error while parsing")
	}
	vendor := models.GetVendorById(ID)
	DeletedVendor := models.DeleteVendorById(vendor)
	res, _ := json.Marshal(DeletedVendor)
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

	//update vendor function
	vendorDetails := models.GetVendorById(ID)

	res, _ := json.Marshal(vendorDetails)
	w.Header().Set("Content-Typer", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetAllAvailibleVendors(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	day := r.URL.Query().Get("day")
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	fmt.Println(day, end, start)
	availableVendors := models.GetAllAvailibleVendors(day, start, end)
	res, _ := json.Marshal(availableVendors)
	w.Header().Set("Content-Typer", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetAllAvailabilityEntries(w http.ResponseWriter, r *http.Request) {
	availabilities, _ := models.GetAllAvailabilityEntries()
	res, _ := json.Marshal(availabilities)
	w.Header().Set("Content-Typer", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func RequestBooking(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	dateString := r.URL.Query().Get("day")
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	vendorID := r.URL.Query().Get("vendorID")
	price := r.URL.Query().Get("price")
	location := r.URL.Query().Get("location")
	vID, er := strconv.Atoi(vendorID)
	if er != nil {
		fmt.Println(er)
	}
	end_time, err := strconv.Atoi(end)
	if err != nil {
		fmt.Print(err)
	}

	start_time, err := strconv.Atoi(start)
	if err != nil {
		fmt.Print(err)
	}
	cost, err := strconv.Atoi(price)
	if err != nil {
		fmt.Print(err)
	}
	booking := models.requestBooking(start_time, end_time, dateString, vID, user, location, cost)
	res, _ := json.Marshal(booking)
	w.Header().Set("Content-Typer", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
