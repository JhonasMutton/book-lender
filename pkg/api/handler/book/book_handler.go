package book

import (
	"encoding/json"
	"github.com/JhonasMutton/book-lender/pkg/api/render"
	"github.com/JhonasMutton/book-lender/pkg/errors"
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
	var bookDTO model.BookDTO
	if err := json.NewDecoder(r.Body).Decode(&bookDTO); err != nil {
		render.ResponseError(w, err, GenerateHTTPErrorStatusCode(errors.ErrInvalidPayload))
		return
	}

	u, err := h.bookUseCase.Create(bookDTO)
	if err != nil {
		render.ResponseError(w, err, GenerateHTTPErrorStatusCode(err))
		return
	}

	render.Response(w, u, http.StatusOK)
}

func GenerateHTTPErrorStatusCode(err error) int {
	switch errors.Cause(err).(type) {
	case *errors.NotFound:
		return http.StatusNotFound
	case *errors.InvalidPayload:
		return http.StatusPreconditionFailed
	case *errors.BadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}