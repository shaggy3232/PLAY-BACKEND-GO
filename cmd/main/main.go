package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/routes"
)

// App contains all components for the application
type App struct {
	API *http.Server
}

func main() {
	// signal handlers
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	app := App{}
	app.Run()

	<-ctx.Done()
	cancel()

	// create a context with timeout for shutting down
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	app.API.Shutdown(ctx)

	log.Print("Application shut down")
}

// Run starts the application components
func (a *App) Run() {
	r := mux.NewRouter()
	routes.RegisterVendorRoutes(r)
	http.Handle("/", r)

	a.API = &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		err := a.API.ListenAndServe()
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatal(err, "failed to start http server")
			}
		}
	}()

	log.Print("Application running")
}
