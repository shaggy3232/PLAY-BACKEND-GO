package controllers

import (
	"context"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

type AvailbilityRepository interface {
	// define crud functions
	CreateAvailbility(ctx context.Context, Availbility models.Availbility) (*models.Availbility, error)
	GetAvailbilitys(ctx context.Context) ([]models.Availbility, error)
	GetAvailbilityById(ctx context.Context, id string) (*models.Availbility, error)
	// TODO: GetFilteredAvailbility()
	DeleteAvailbility(ctx context.Context, id string) (int, error)
}

type AvailbilityController struct {
	Store AvailbilityRepository
}

func (c *AvailbilityController) GetAvailbilityById(ctx context.Context, id string) (*models.Availbility, error) {
	Availbility, err := c.Store.GetAvailbilityById(ctx, id)

	if err != nil {
		return nil, err
	}

	return Availbility, nil

}

func (c *AvailbilityController) GetAvailbilitys(ctx context.Context) ([]models.Availbility, error) {
	Availbilitys, err := c.Store.GetAvailbilitys(ctx)

	if err != nil {
		return nil, err
	}

	return Availbilitys, nil

}

func (c *AvailbilityController) CreateAvailbility(ctx context.Context, newAvailbility *models.Availbility) (*models.Availbility, error) {
	createdAvailbility, err := c.Store.CreateAvailbility(ctx, *newAvailbility)
	if err != nil {
		return nil, err
	}

	return createdAvailbility, nil
}

func (c *AvailbilityController) DeleteAvailbility(ctx context.Context, id string) (*models.Availbility, error) {
	return nil, nil

}
