package api

import (
	"go-swag/internal/handler"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users/signup", handler.SignupHandler)
	mux.HandleFunc("/users", handler.GetAllUsersHandler)
	mux.HandleFunc("/users/{id}", handler.GetUserByIDHandler)
	mux.HandleFunc("/users/login", handler.LoginHandler)
}