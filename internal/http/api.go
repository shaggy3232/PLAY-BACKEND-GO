package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/controllers"
	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/http/middleware"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Port                   int
	Server                 *http.Server
	UserController         *controllers.UserController
	AvailabilityController *controllers.AvailabilityController
	BookingController      *controllers.BookingController
}

// APIServerOption defines a function that modifies the Server.
type APIServerOption func(*APIServer)

type APIError struct {
	Message string `json:"message"`
}

// NewAPIServer setups a the API + Routes
func NewAPIServer(options ...APIServerOption) *APIServer {
	api := &APIServer{
		Port: 8080,
	}

	// Apply the provided options
	for _, option := range options {
		option(api)
	}

	r := mux.NewRouter()

	// middleware
	r.Use(middleware.NewPanicMiddleware())
	r.Use(middleware.NewLoggingMiddleware())
	r.Use(middleware.CORSMiddleware())

	protectedRoutes := r.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.JWTMiddleware)

	//Login Routes
	r.HandleFunc("/login", api.HandleUserLogin).Methods("POST", "OPTIONS")
	// User routes
	r.HandleFunc("/users", api.HandleCreateUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/users", api.HandleUpdateUser).Methods("PUT", "OPTIONS")
	r.HandleFunc("/users/{userID}", api.HandleGetUserById).Methods("GET", "OPTIONS")
	r.HandleFunc("/users", api.HandleListUsers).Methods("GET", "OPTIONS")
	r.HandleFunc("/users/{start}/{end}", api.HandleGetAvailableUsers).Methods("GET", "OPTIONS")
	r.HandleFunc("/users/{userID}", api.HandleDeleteUser).Methods("DELETE", "OPTIONS")

	// Availability routes
	r.HandleFunc("/availabilities", api.HandleCreateAvailability).Methods("POST", "OPTIONS")
	r.HandleFunc("/availabilities", api.handleUpdateAvailability).Methods("PUT", "OPTIONS")
	r.HandleFunc("/availabilities/{availabilityID}", api.HandleDeleteAvailability).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/availabilities/{availabilityID}", api.HandleGetAvailabilityById).Methods("GET", "OPTIONS")
	r.HandleFunc("/availabilities/user/{userID}", api.HandleGetUsersAvailability).Methods("GET", "OPTIONS")
	r.HandleFunc("/availabilities/{start}/{end}", api.HandleGetValidAvailabilities).Methods("GET", "OPTIONS")
	r.HandleFunc("/availabilities", api.HandleListAvailabilities).Methods("GET", "OPTIONS")

	// Bookings routes
	r.HandleFunc("/bookings", api.HandleCreateBooking).Methods("POST", "OPTIONS")
	r.HandleFunc("/bookings", api.HandleEditBookings).Methods("PUT", "OPTIONS")
	r.HandleFunc("/bookings/{bookingID}", api.handleAcceptBooking).Methods("PUT", "OPTIONS")
	r.HandleFunc("/bookings/{bookingID}", api.HandleGetBookingById).Methods("GET", "OPTIONS")
	r.HandleFunc("/bookings/referee/{refereeID}", api.HandleGetBookingByRef).Methods("GET", "OPTIONS")
	r.HandleFunc("/bookings/user/{userID}", api.HandleGetBookingByUser).Methods("GET", "OPTIONS")
	r.HandleFunc("/bookings", api.HandleListBookings).Methods("GET", "OPTIONS")
	r.HandleFunc("/bookings/{bookingID}", api.HandleDeleteBooking).Methods("DELETE", "OPTIONS")

	http.Handle("/users", r)

	http.Handle("/", r)

	api.Server = &http.Server{
		Addr:    fmt.Sprintf(":%d", api.Port),
		Handler: r,
	}

	return api
}

func WithPort(port int) APIServerOption {
	return func(a *APIServer) {
		a.Port = port
	}
}

func WithUserController(userController *controllers.UserController) APIServerOption {
	return func(a *APIServer) {
		a.UserController = userController
	}
}
func WithBookingController(bookingController *controllers.BookingController) APIServerOption {
	return func(a *APIServer) {
		a.BookingController = bookingController
	}
}

func WithAvailabilityController(availabilityController *controllers.AvailabilityController) APIServerOption {
	return func(a *APIServer) {
		a.AvailabilityController = availabilityController
	}
}

func (a *APIServer) Shutdown(ctx context.Context) error {
	return a.Server.Shutdown(ctx)
}
