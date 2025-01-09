package controllers

import (
	"context"
	"net/http"

	playhttp "github.com/shaggy3232/PLAY-BACKEND-GO/pkg/http"
	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/models"
	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/repositories"
	"github.com/shaggy3232/PLAY-BACKEND-GO/pkg/utils"
)

type UserController struct {
	Repo repositories.UserRepository
}

//handle http request using the repository that is passed in to the controller

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	potentialUser := &models.User{}
	ctx := context.TODO()
	utils.ParseBody(r, potentialUser)
	newUser, err := c.Repo.CreateUser(ctx, *potentialUser)
	if err != nil {
		println(err)
	}

	playhttp.Encode(w, r, http.StatusOK, newUser)

}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {

}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

}
