package controllers

import (
	"context"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

type UserRepository interface {
	// define crud functions
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	GetUsers(ctx context.Context) (*models.UserList, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	// TODO: GetFilteredUser()
	UpdateUser(ctx context.Context, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id string) (int, error)
}

type UserController struct {
	store UserRepository
}

func (c *UserController) GetUserById(ctx context.Context, id string) (*models.User, error) {
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

func (c *UserController) DeleteUser(ctx context.Context, id string) (*models.User, error) {
	return nil, nil

}

func (c *UserController) UpdateUser(ctx context.Context, newUser models.User) (*models.User, error) {
	return nil, nil
}
