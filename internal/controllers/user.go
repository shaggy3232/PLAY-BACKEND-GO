package controllers

import (
	"context"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/repositories"
)

type UserController struct {
	store repositories.UserRepository
}

func (c *UserController) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	user, err := c.store.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (c *UserController) GetUsers(ctx context.Context) (*models.UserList, error) {
	users, err := c.store.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil

}

func (c *UserController) CreateUser(ctx context.Context, newUser *models.User) (*models.User, error) {
	createdUser, err := c.store.CreateUser(ctx, *newUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (c *UserController) DeleteUser(ctx context.Context, id int) (*models.User, error) {
	return nil, nil

}

func (c *UserController) UpdateUser(ctx context.Context, newUser models.User) (*models.User, error) {
	return nil, nil
}
