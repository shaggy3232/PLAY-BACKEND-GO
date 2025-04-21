package http

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/http/auth"
	"github.com/shaggy3232/PLAY-BACKEND-GO/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// user requests

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
	//create a token and assign it in the response as a cookie
	token, err := auth.GenerateJWT(newUser.ID)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Could not generate a token")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "Could not generate Token"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
	})

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

func (api *APIServer) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	newUser, err := decode[models.User](r)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to create user")
	}

	updatedUser, err := api.UserController.UpdateUser(r.Context(), &newUser)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to create user")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "internal server error"})
		return
	}
	if err := encode(w, r, http.StatusOK, updatedUser); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode the updatedUser json response")
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
		return
	}

	user, err := api.UserController.GetUserById(r.Context(), userId)
	if err != nil {
		// TODO: distinguish between missing users and actual errors
		log.Error().
			Err(err).
			Msg("failed to get user")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get user"})
		return
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
		log.Error().
			Err(err).
			Msg("failed to delete user")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get user"})
		return
	}

	if err := encode(w, r, http.StatusOK, deletedUser); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode deleted User json response")
	}
}

func (api *APIServer) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	creds, err := decode[models.Login](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, &APIError{Message: "invalid login format"})
		return
	}
	//validate that the password is correct and and the user exist
	user, err := api.UserController.GetUserFromEmail(r.Context(), creds.Email)

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to get user")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get user"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))

	if err != nil {
		log.Error().
			Err(err).
			Msg("Passwords do not match")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "passwords do not match"})
		return
	}

	//create a token and assign it in the response as a cookie
	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Could not generate a token")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "Could not generate Token"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,                    // Prevents JavaScript access (protects against XSS)
		Secure:   true,                    // Requires HTTPS
		SameSite: http.SameSiteStrictMode, // Prevents CSRF
		Expires:  time.Now().Add(1 * time.Hour),
	})
	w.WriteHeader(http.StatusOK)

}

func (api *APIServer) HandleGetAvailableUsers(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	startString, ok := vars["start"]

	if !ok {
		log.Error().Msg("Failed to get startTime from request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing start time in request"})
	}

	endString, ok := vars["end"]

	if !ok {
		log.Error().Msg("Failed to get end time from the request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing end time in request"})
		return
	}

	start, err := time.Parse(time.RFC3339, startString)
	if err != nil {
		log.Error().Msg("Failed to parse start time in to time.Time")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "bad start time format"})
		return
	}
	end, err := time.Parse(time.RFC3339, endString)
	if err != nil {
		log.Error().Msg("failed to parse end time in to time.Time")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "bad end time format"})
		return
	}

	users, err := api.UserController.GetAvailableUsers(r.Context(), start, end)
	if err != nil {
		log.Error().Err(err).Msg("failed to get avavilable users")
		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get available users"})
	}

	if err := encode(w, r, http.StatusOK, users); err != nil {
		log.Error().Err(err).Msg("Failed to encode the available users json response")
	}
}

// booking requests

func (api *APIServer) HandleCreateBooking(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	potentialBooking, err := decode[models.Booking](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, &APIError{Message: "invalid booking schema"})
		return
	}

	newBooking, err := api.BookingController.CreateBooking(r.Context(), &potentialBooking)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to create booking")
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
			Msg("failed to get bookings")
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
	bookingID, ok := vars["bookingID"]
	if !ok {
		log.Error().
			Msg("failed to get id from request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing user id in request"})
		return
	}

	user, err := api.BookingController.GetBookingById(r.Context(), bookingID)
	if err != nil {
		// TODO: distinguish between missing bookings and actual errors
		log.Error().
			Err(err).
			Msg("failed to get booking")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get user"})
		return
	}

	if err := encode(w, r, http.StatusOK, user); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode GetBookings json response")
	}
}

func (api *APIServer) HandleGetBookingByRef(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	refereeID, ok := vars["refereeID"]
	if !ok {
		log.Error().
			Msg("failed to get id from request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing user id in request"})
		return
	}
	bookings, err := api.BookingController.GetBookingsByRef(r.Context(), refereeID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get referee's Booking")
		encode(w, r, http.StatusInternalServerError, &APIError{Message: "Failed to get referee's Bookings"})
	}
	if err := encode(w, r, http.StatusOK, bookings); err != nil {
		log.Error().Err(err).Msg("failed encodeing bookings to json response")
	}
}
func (api *APIServer) HandleGetBookingByUser(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	userID, ok := vars["userID"]
	if !ok {
		log.Error().
			Msg("failed to get id from request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing user id in request"})
		return
	}
	bookings, err := api.BookingController.GetBookingsByUser(r.Context(), userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get referee's Booking")
		encode(w, r, http.StatusInternalServerError, &APIError{Message: "Failed to get referee's Bookings"})
	}
	if err := encode(w, r, http.StatusOK, bookings); err != nil {
		log.Error().Err(err).Msg("failed encodeing bookings to json response")
	}
}

func (api *APIServer) HandleDeleteBooking(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	bookingID, ok := vars["bookingID"]
	if !ok {
		log.Error().
			Msg("failed to get id from request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing user id in request"})
		return
	}

	deletedBooking, err := api.BookingController.DeleteBooking(r.Context(), bookingID)
	if err != nil {
		// TODO: distinguish between missing users and actual errors
		log.Error().
			Err(err).
			Msg("failed to delete booking")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get user"})
		return
	}

	if err := encode(w, r, http.StatusOK, deletedBooking); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode deleted Booking json response")
	}
}

func (api *APIServer) HandleEditBookings(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	booking, err := decode[models.Booking](r)
	if err != nil {
		log.Error().Err(err).Msg("could not decode request into booking")
	}

	updatedBooking, err := api.BookingController.EditBooking(r.Context(), booking)

	if err != nil {
		log.Error().Err(err).Msg("Could not update the booking")
		encode(w, r, http.StatusInternalServerError, &APIError{Message: "Could not update the booking"})
	}

	if err := encode(w, r, http.StatusOK, updatedBooking); err != nil {
		log.Error().Err(err).Msg("failed to create Json response with the updated booking")
	}

}

func (api *APIServer) handleAcceptBooking(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	bookingID, ok := vars["bookingID"]
	if !ok {
		log.Error().Msg("failed to get booking id from request")
		encode(w, r, http.StatusBadRequest, &APIError{
			Message: "COULD NOT FIND BOOKING ID IN THE REQUEST",
		})
	}

	booking, err := api.BookingController.AcceptBooking(r.Context(), bookingID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to accept booking")
		encode(w, r, http.StatusInternalServerError, &APIError{
			Message: "COULD NOT ACCEPT BOOKING",
		})
		return
	}

	if err := encode(w, r, http.StatusOK, booking); err != nil {
		log.Error().Err(err).Msg("could not encode booking into json")
		return
	}

}

//availability requests

func (api *APIServer) HandleCreateAvailability(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	potentialAvailability, err := decode[models.Availability](r)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Cannot decode the request into availability model")
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

func (api *APIServer) HandleListAvailabilities(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	availabilities, err := api.AvailabilityController.GetAvailabilities(r.Context())

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to get availabilities")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "internal server error"})
		return
	}

	if err := encode[[]models.Availability](w, r, http.StatusOK, availabilities); err != nil {
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
		return
	}

	availability, err := api.AvailabilityController.GetAvailabilityById(r.Context(), availabilityId)
	if err != nil {
		// TODO: distinguish between missing availabilities and actual errors
		log.Error().
			Err(err).
			Msg("failed to get availability")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get availability"})
		return
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
	log.Print(deletedAvailability)
	if err != nil {
		// TODO: distinguish between missing availabilities and actual errors
		log.Error().
			Err(err).
			Msg("failed to delete availability")

		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get availability"})
		return
	}

	if err := encode(w, r, http.StatusOK, deletedAvailability); err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode deleted Availability json response")
		return
	}
}

func (api *APIServer) HandleGetUsersAvailability(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	userID, ok := vars["userID"]
	if !ok {
		log.Error().
			Msg("failed to get id from request")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "missing availability id in request"})
		return
	}

	userAvailability, err := api.AvailabilityController.GetAvailabilityByUser(r.Context(), userID)
	if err != nil {
		log.Error().
			Msg("Failed to get availability from the database")
		encode(w, r, http.StatusInternalServerError, &APIError{Message: "failed to get availabilities"})
		return
	}
	if err := encode(w, r, http.StatusOK, userAvailability); err != nil {
		log.Error().Msg("FAILED TO ENCODE THE AVAILABILITIES TO JSON")
	}

}

func (api *APIServer) handleUpdateAvailability(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	avail, err := decode[models.Availability](r)

	if err != nil {
		log.Error().Err(err).Msg("Could not decode request into avaialbility")
		encode(w, r, http.StatusBadRequest, &APIError{Message: "invalid availability schema"})

	}
	updatedAvail, err := api.AvailabilityController.UpdateAvailability(r.Context(), avail)
	if err != nil {
		log.Error().Err(err).Msg("could not update the availability")
		encode(w, r, http.StatusInternalServerError, &APIError{Message: "could not update the availability"})

	}

	if err := encode(w, r, http.StatusOK, updatedAvail); err != nil {
		log.Error().Err(err).Msg("could not encould updated availability into JSON")
	}
}
