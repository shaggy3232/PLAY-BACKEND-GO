package controllers

import (
	"context"
	"net/http"

	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

type UserStore interface {
	// define crud functions
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	GetUserById(ctx context.Context, id int) (*models.User, error)
	// TODO: GetFilteredUser()
	UpdateUser(ctx context.Context, id int, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int) (int, error)
}

type UserController struct {
	store UserStore
}

//handle http request using the repository that is passed in to the controller

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {

}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {

}

func (c *UserController) CreateUser(ctx context.Context, newUser *models.User) (*models.User, error) {
	createdUser, err := c.store.CreateUser(ctx, *newUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

}
