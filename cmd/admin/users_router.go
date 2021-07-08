package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (a *AdminAPI) UsersRouter(router *mux.Router) {
	usersHandler := NewUsersHandler(a.userRepository)

	router.
		Methods(http.MethodGet).
		Path("/users").
		HandlerFunc(usersHandler.Get)

	router.
		Methods(http.MethodGet).
		Path("/users/{id}").
		HandlerFunc(usersHandler.GetByID)
}
