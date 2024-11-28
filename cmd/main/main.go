package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/http/middleware"
	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/routes"
)

// App contains all components for the application
type App struct {
	API *http.Server
}

func main() {
	// signal handlers
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// create the application and start it up
	app := App{}
	app.Run()

	<-ctx.Done()
	cancel()

	// shutdown the applicaiton
	app.Stop()
}

// Run starts the application components
func (a *App) Run() {
	// setup logging
	// TODO: add log_level as config var
	logLevel := zerolog.DebugLevel
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log := zerolog.New(os.Stderr).
		Level(logLevel).
		With().
		Timestamp().
		Logger()
	zerolog.DefaultContextLogger = &log

	r := mux.NewRouter()

	// middleware
	r.Use(middleware.NewPanicMiddleware())
	r.Use(middleware.NewLoggingMiddleware())

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
				log.Fatal().
					Err(err).
					Msg("Failed to start http server")
			}
		}
	}()

	log.Info().Msg("Application running")
}

// Stop gracefully shutsdown the application
func (a *App) Stop() {
	// create a context with timeout for shutting down
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	a.API.Shutdown(ctx)

	log.Info().Msg("Application shut down")
}
