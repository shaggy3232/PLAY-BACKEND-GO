package controllers

import (
	"context"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

type AvailabilityRepository interface {
	// define crud functions
	CreateAvailability(ctx context.Context, Availability models.Availability) (*models.Availability, error)
	GetAvailabilities(ctx context.Context) ([]models.Availability, error)
	GetAvailabilityById(ctx context.Context, id string) (*models.Availability, error)
	// TODO: GetFilteredAvailability()
	DeleteAvailability(ctx context.Context, id string) (int, error)
}

type AvailabilityController struct {
	Store AvailabilityRepository
}

func (c *AvailabilityController) GetAvailabilityById(ctx context.Context, id string) (*models.Availability, error) {
	availability, err := c.Store.GetAvailabilityById(ctx, id)

	if err != nil {
		return nil, err
	}

	return availability, nil
}

func (c *AvailabilityController) GetAvailabilities(ctx context.Context) ([]models.Availability, error) {
	availabilities, err := c.Store.GetAvailabilities(ctx)

	if err != nil {
		return nil, err
	}

	return availabilities, nil
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
