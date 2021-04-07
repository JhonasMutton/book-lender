package book

import (
	"encoding/json"
	"github.com/JhonasMutton/book-lender/pkg/api/render"
	"github.com/JhonasMutton/book-lender/pkg/log"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/usecase/book"
	"net/http"
)

type Handler struct {
	bookUseCase book.IUseCase
}

func NewHandler(bookUseCase book.IUseCase) *Handler {
	return &Handler{bookUseCase: bookUseCase}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("Handling book post!")
	var bookDTO model.BookDTO
	if err := json.NewDecoder(r.Body).Decode(&bookDTO); err != nil {
		log.Logger.Errorf("Error to decode book, error: %s", err.Error())
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	u, err := h.bookUseCase.Create(bookDTO)
	if err != nil {
		log.Logger.Errorf("Some error has occurred on Create, error: %s", err.Error())
		render.ResponseError(w, err, render.GenerateHTTPErrorStatusCode(err))
		return
	}

	render.Response(w, u, http.StatusCreated)
	log.Logger.Info("Post was handled")
}
