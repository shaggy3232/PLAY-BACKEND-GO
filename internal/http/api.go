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
	Port           int
	Server         *http.Server
	UserController *controllers.UserController
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

	// routes
	r.HandleFunc("/users", api.HandleCreateUser).Methods("POST")
	r.HandleFunc("/users", api.HandleListUsers).Methods("GET")
	r.HandleFunc("/users/{userID}", api.HandleGetUserById).Methods("GET")
	r.HandleFunc("/users/{userID}", api.HandleDeleteUser).Methods("DELETE")

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

func (a *APIServer) Shutdown(ctx context.Context) error {
	return a.Server.Shutdown(ctx)
}
