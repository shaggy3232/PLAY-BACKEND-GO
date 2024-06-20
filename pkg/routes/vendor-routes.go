package routes

import (
	"github.com/gorilla/mux"
	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/controllers"
)

var RegisterVendorRoutes = func(router *mux.Router) {
	router.HandleFunc("/vendor/", controllers.CreateVendor).Methods("POST")
	router.HandleFunc("/vendor/", controllers.GetVendor).Methods("GET")
	router.HandleFunc("/vendor/{vendorId}", controllers.GetVendorById).Methods("GET")
	router.HandleFunc("/vendor/{vendorId}", controllers.UpdateVendor).Methods("PUT")
	router.HandleFunc("/vendor/{vendorId}", controllers.DeleteVendor).Methods("DELETE")
}
