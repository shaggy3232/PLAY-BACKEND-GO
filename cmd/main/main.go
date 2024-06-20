package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterVendorRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
