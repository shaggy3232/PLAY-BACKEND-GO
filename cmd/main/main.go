package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/controllers"
	playhttp "github.com/shaggy3232/PLAY-BACKEND-GO/internal/http"
	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/postgres"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// App contains all components for the application
type App struct {
	API *playhttp.APIServer
	DB  *postgres.Client
}

func main() {
	// signal handlers
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// create the application and start it up
	app := App{}
	app.Run(ctx)

	<-ctx.Done()
	cancel()

	// shutdown the application
	app.Stop()
}

// Run starts the application components
func (a *App) Run(ctx context.Context) {
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

	// TODO: Make config struct with env variable parsing
	db, err := postgres.New(
		ctx,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to connect to postgres DB")
	}

	userController := controllers.UserController{Store: db}
	availabilityController := controllers.AvailabilityController{Store: db}
	bookingController := controllers.BookingController{Store: db}

	api := playhttp.NewAPIServer(
		playhttp.WithPort(8080),
		playhttp.WithUserController(&userController),
		playhttp.WithAvailabilityController(&availabilityController),
		playhttp.WithBookingController(&bookingController),
	)

	a.API = api
	a.DB = db

	go func() {
		err := a.API.Server.ListenAndServe()
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

// Stop gracefully shuts down the application
func (a *App) Stop() {
	// create a context with timeout for shutting down
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	a.API.Shutdown(ctx)
	a.DB.Shutdown()

	log.Info().Msg("Application shut down")
}
