package config

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/users", nil).Methods(http.MethodPost)

	return r
}
