package http

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
)

func (api *APIServer) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	potentialUser, err := decode[models.User](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, &APIError{Message: "invalid user schema"})
		return
	}
	newUser, err := api.UserController.CreateUser(r.Context(), &potentialUser)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to create user")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "internal server error"})
		return
	}

	if err := encode(w, r, http.StatusOK, newUser); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode CreateUser json response")
	}
}

func (api *APIServer) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	users, err := api.UserController.GetUsers(r.Context())

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to create user")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "internal server error"})
		return
	}

	if err := encode(w, r, http.StatusOK, users); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode GetUsers json response")
	}

}

func (api *APIServer) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	userId := vars["user_ID"]
	ID, err := strconv.ParseInt(userId, 0, 0)

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to get id from request")
	}

	user, err := api.UserController.GetUserById(r.Context(), ID)

	if err := encode(w, r, http.StatusOK, user); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode GetUsers json response")
	}

}

func (api *APIServer) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	potentialUser, err := decode[models.User](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, &APIError{Message: "invalid user schema"})
		return
	}

	updatedUser, err := api.UserController.UpdateUser(r.Context(), potentialUser)

	if err := encode(w, r, http.StatusOK, updatedUser); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode GetUsers json response")
	}

}

func (api *APIServer) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	userId := vars["user_ID"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to get id from request")
	}

	deletedUser, err := api.UserController.DeleteUser(r.Context(), int(ID))

	if err := encode(w, r, http.StatusOK, deletedUser); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode deleted User json response")
	}
}
