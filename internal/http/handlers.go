package http

import (
	"net/http"

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

func (api *APIServer) HandleListUsers(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	users, err := api.UserController.GetUsers(r.Context())

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to create user")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "internal server error"})
		return
	}

	if err := encode[[]models.User](w, r, http.StatusOK, users); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode GetUsers json response")
	}

}

func (api *APIServer) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	userId, ok := vars["userID"]
	if !ok {
		log.Error().
			Msg("failed to get id from request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing user id in request"})
	}

	user, err := api.UserController.GetUserById(r.Context(), userId)
	if err != nil {
		// TODO: distinguish between missing users and actual errors
		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get user"})
	}

	if err := encode(w, r, http.StatusOK, user); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode GetUsers json response")
	}
}

func (api *APIServer) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	userId, ok := vars["userID"]
	if !ok {
		log.Error().
			Msg("failed to get id from request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing user id in request"})
	}

	deletedUser, err := api.UserController.DeleteUser(r.Context(), userId)
	log.Print(deletedUser)
	if err != nil {
		// TODO: distinguish between missing users and actual errors
		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get user"})
	}

	if err := encode(w, r, http.StatusOK, deletedUser); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode deleted User json response")
	}
}

func (api *APIServer) HandleCreateBooking(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	potentialBooking, err := decode[models.Booking](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, &APIError{Message: "invalid user schema"})
		return
	}
	newBooking, err := api.BookingController.CreateBooking(r.Context(), &potentialBooking)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to create user")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "internal server error"})
		return
	}

	if err := encode(w, r, http.StatusOK, newBooking); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode CreateBooking json response")
	}
}

func (api *APIServer) HandleListBookings(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	users, err := api.BookingController.GetBookings(r.Context())

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to create user")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "internal server error"})
		return
	}

	if err := encode[[]models.Booking](w, r, http.StatusOK, users); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode GetBookings json response")
	}

}

func (api *APIServer) HandleGetBookingById(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	userId, ok := vars["userID"]
	if !ok {
		log.Error().
			Msg("failed to get id from request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing user id in request"})
	}

	user, err := api.BookingController.GetBookingById(r.Context(), userId)
	if err != nil {
		// TODO: distinguish between missing users and actual errors
		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get user"})
	}

	if err := encode(w, r, http.StatusOK, user); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode GetBookings json response")
	}
}

func (api *APIServer) HandleDeleteBooking(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	userId, ok := vars["userID"]
	if !ok {
		log.Error().
			Msg("failed to get id from request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing user id in request"})
	}

	deletedBooking, err := api.BookingController.DeleteBooking(r.Context(), userId)
	if err != nil {
		// TODO: distinguish between missing users and actual errors
		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get user"})
	}

	if err := encode(w, r, http.StatusOK, deletedBooking); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode deleted Booking json response")
	}
}

func (api *APIServer) HandleCreateAvailability(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	potentialAvailability, err := decode[models.Availability](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, &APIError{Message: "invalid availability schema"})
		return
	}
	newAvailability, err := api.AvailabilityController.CreateAvailability(r.Context(), &potentialAvailability)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to create availability")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "internal server error"})
		return
	}

	if err := encode(w, r, http.StatusOK, newAvailability); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode CreateAvailability json response")
	}
}

func (api *APIServer) HandleListAvailabilitys(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	availabilitys, err := api.AvailabilityController.GetAvailabilitys(r.Context())

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to create availability")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "internal server error"})
		return
	}

	if err := encode[[]models.Availability](w, r, http.StatusOK, availabilitys); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode GetAvailabilitys json response")
	}

}

func (api *APIServer) HandleGetAvailabilityById(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	availabilityId, ok := vars["availabilityID"]
	if !ok {
		log.Error().
			Msg("failed to get id from request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing availability id in request"})
	}

	availability, err := api.AvailabilityController.GetAvailabilityById(r.Context(), availabilityId)
	if err != nil {
		// TODO: distinguish between missing availabilitys and actual errors
		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get availability"})
	}

	if err := encode(w, r, http.StatusOK, availability); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode GetAvailabilitys json response")
	}
}

func (api *APIServer) HandleDeleteAvailability(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	availabilityId, ok := vars["availabilityID"]
	if !ok {
		log.Error().
			Msg("failed to get id from request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing availability id in request"})
	}

	deletedAvailability, err := api.AvailabilityController.DeleteAvailability(r.Context(), availabilityId)
	if err != nil {
		// TODO: distinguish between missing availabilitys and actual errors
		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get availability"})
	}

	if err := encode(w, r, http.StatusOK, deletedAvailability); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode deleted Availability json response")
	}
}
