package controllers

import (
	"context"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

type AvailabilityRepository interface {
	// define crud functions
	CreateAvailability(ctx context.Context, Availability models.Availability) (*models.Availability, error)
	GetAvailabilitys(ctx context.Context) ([]models.Availability, error)
	GetAvailabilityById(ctx context.Context, id string) (*models.Availability, error)
	// TODO: GetFilteredAvailability()
	DeleteAvailability(ctx context.Context, id string) (int, error)
}

type AvailabilityController struct {
	Store AvailabilityRepository
}

func (c *AvailabilityController) GetAvailabilityById(ctx context.Context, id string) (*models.Availability, error) {
	Availability, err := c.Store.GetAvailabilityById(ctx, id)

	if err != nil {
		return nil, err
	}

	return Availability, nil

}

func (c *AvailabilityController) GetAvailabilitys(ctx context.Context) ([]models.Availability, error) {
	Availabilitys, err := c.Store.GetAvailabilitys(ctx)

	if err != nil {
		return nil, err
	}

	return Availabilitys, nil

}

func (c *AvailabilityController) CreateAvailability(ctx context.Context, newAvailability *models.Availability) (*models.Availability, error) {
	createdAvailability, err := c.Store.CreateAvailability(ctx, *newAvailability)
	if err != nil {
		return nil, err
	}

	return createdAvailability, nil
}

func (c *AvailabilityController) DeleteAvailability(ctx context.Context, id string) (*models.Availability, error) {
	return nil, nil

}
