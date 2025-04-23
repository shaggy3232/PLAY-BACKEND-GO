package controllers

import (
	"context"
	"time"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

type AvailabilityRepository interface {
	// define crud functions
	CreateAvailability(ctx context.Context, Availability *models.Availability) (*models.Availability, error)
	GetAvailabilities(ctx context.Context) ([]models.Availability, error)
	GetAvailabilityById(ctx context.Context, id string) (*models.Availability, error)
	// TODO: GetFilteredAvailability()
	GetAvailabilityByUser(ctx context.Context, userId string) ([]models.Availability, error)
	UpdateAvailability(ctx context.Context, avaialbility models.Availability) (*models.Availability, error)
	DeleteAvailability(ctx context.Context, id string) (*models.Availability, error)
	GetValidAvailabilities(ctx context.Context, start time.Time, end time.Time) ([]models.Availability, error)
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
	createdAvailability, err := c.Store.CreateAvailability(ctx, newAvailability)
	if err != nil {
		return nil, err
	}

	return createdAvailability, nil
}

func (c *AvailabilityController) DeleteAvailability(ctx context.Context, id string) (*models.Availability, error) {
	deletedAvailability, err := c.Store.DeleteAvailability(ctx, id)

	if err != nil {
		return nil, err
	}

	return deletedAvailability, nil

}

func (c *AvailabilityController) GetAvailabilityByUser(ctx context.Context, id string) ([]models.Availability, error) {
	userAvailability, err := c.Store.GetAvailabilityByUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return userAvailability, nil
}

func (c *AvailabilityController) UpdateAvailability(ctx context.Context, availability models.Availability) (*models.Availability, error) {
	updatedAvail, err := c.Store.UpdateAvailability(ctx, availability)
	if err != nil {
		return nil, err
	}
	return updatedAvail, nil
}

func (c *AvailabilityController) GetValidAvailabilities(ctx context.Context, start time.Time, end time.Time) ([]models.Availability, error) {
	availabilities, err := c.Store.GetValidAvailabilities(ctx, start, end)

	if err != nil {
		return nil, err
	}
	return availabilities, nil
}
