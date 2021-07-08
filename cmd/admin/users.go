package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	domain "multiapi"
	"net/http"
	httputil "multiapi/pkg/http"
)

type CreateUserPayload struct {
	Name  string
	Email string
}

func (p *CreateUserPayload) validate() []string {
	var errs []string

	if len(p.Name) < 3 {
		errs = append(errs, "name: invalid length")
	}

	if len(p.Email) < 5 {
		errs = append(errs, "email: invalid length")
	}

	return errs
}

// UsersHandler represents the set of handlers for the users API.
type UsersHandler struct {
	userRepo  domain.UserRepository
}

// NewUsersHandler returns an initialized users handler with the required dependencies.
func NewUsersHandler(userRepo domain.UserRepository) *UsersHandler {
	return &UsersHandler{
		userRepo:  userRepo,
	}
}

// Get gets all users
func (h *UsersHandler) Get(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepo.GetAll(r.Context())
	if err != nil {
		httputil.RespondInternalError(w)
		return
	}

	httputil.RespondJSON(w, users, http.StatusOK)
}

// GetByID tries to get a user by ID
func (h *UsersHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, ok := vars["id"]
	if !ok {
		log.Error("could not read id param in UsersHandler.GetByID")
		httputil.RespondInternalError(w)
		return
	}

	user, err := h.userRepo.FindByID(r.Context(), uid)
	if err != nil {
		httputil.RespondInternalError(w)
		return
	}

	if user == nil {
		httputil.RespondJSON(w, &httputil.AppError{
			Errors: []string{"user not found"},
		}, http.StatusNotFound)
		return
	}

	httputil.RespondJSON(w, user, http.StatusOK)
}
