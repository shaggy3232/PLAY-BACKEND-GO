package controllers

import (
	"context"
	"time"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

type BookingRepository interface {
	// define crud functions
	CreateBooking(ctx context.Context, Booking *models.Booking) (*models.Booking, error)
	GetBookings(ctx context.Context) ([]models.Booking, error)
	GetBookingById(ctx context.Context, id string) (*models.Booking, error)
	// TODO: GetFilteredBooking()
	CheckConflicts(ctx context.Context, userID string, start time.Time, end time.Time) (bool, error)
	DeleteBooking(ctx context.Context, id string) (*models.Booking, error)
}

type BookingController struct {
	Store BookingRepository
}

func (c *BookingController) GetBookingById(ctx context.Context, id string) (*models.Booking, error) {
	Booking, err := c.Store.GetBookingById(ctx, id)

	if err != nil {
		return nil, err
	}

	return Booking, nil

}

func (c *BookingController) GetBookings(ctx context.Context) ([]models.Booking, error) {
	Bookings, err := c.Store.GetBookings(ctx)

	if err != nil {
		return nil, err
	}

	return Bookings, nil

}

func (c *BookingController) CreateBooking(ctx context.Context, newBooking *models.Booking) (*models.Booking, error) {
	createdBooking, err := c.Store.CreateBooking(ctx, newBooking)
	if err != nil {
		return nil, err
	}

	return createdBooking, nil
}

func (c *BookingController) DeleteBooking(ctx context.Context, id string) (*models.Booking, error) {
	return nil, nil

}

func (c *BookingController) CheckConflicts(ctx context.Context, userID string, start time.Time, end time.Time) (bool, error) {
	isConflict, err := c.Store.CheckConflicts(ctx, userID, start, end)
	if err != nil {
		return true, err
	}

	return isConflict, nil

}
