package controllers

import (
	"context"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

type UserRepository interface {
	// define crud functions
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	// TODO: GetFilteredUser()
	DeleteUser(ctx context.Context, id string) (*models.User, error)
}

type UserController struct {
	Store UserRepository
}

func (c *UserController) GetUserById(ctx context.Context, id string) (*models.User, error) {
	user, err := c.Store.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (c *UserController) GetUsers(ctx context.Context) ([]models.User, error) {
	users, err := c.Store.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil

}

func (c *UserController) CreateUser(ctx context.Context, newUser *models.User) (*models.User, error) {
	createdUser, err := c.Store.CreateUser(ctx, *newUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (c *UserController) DeleteUser(ctx context.Context, id string) (*models.User, error) {

	deletedUser, err := c.Store.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return deletedUser, nil

}
