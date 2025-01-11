package http

import (
	"net/http"

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
