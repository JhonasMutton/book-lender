package config

import (
	"github.com/JhonasMutton/book-lender/pkg/api/handler/user"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandlerConfig(userHandler *user.Handler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", userHandler.FindUsers).Methods(http.MethodGet)

	return r
}
