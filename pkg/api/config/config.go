package config

import (
	"github.com/JhonasMutton/book-lender/pkg/api/handler/book"
	"github.com/JhonasMutton/book-lender/pkg/api/handler/lend"
	"github.com/JhonasMutton/book-lender/pkg/api/handler/user"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandlerConfig(userHandler *user.Handler, bookHandler *book.Handler, lendHandler *lend.Handler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/user", userHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/user", userHandler.Find).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", userHandler.FindById).Methods(http.MethodGet)

	r.HandleFunc("/book", bookHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/book/lend", lendHandler.Lend).Methods(http.MethodPut)
	r.HandleFunc("/book/return", lendHandler.Return).Methods(http.MethodPut)

	return r
}
