package config

import (
	"github.com/JhonasMutton/book-lender/pkg/api/handler/user"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandlerConfig(userHandler *user.Handler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/user", userHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/user", userHandler.Find).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", userHandler.FindById).Methods(http.MethodGet)

	return r
}
