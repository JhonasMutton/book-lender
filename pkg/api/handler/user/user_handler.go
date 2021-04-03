package user

import (
	"encoding/json"
	"github.com/JhonasMutton/book-lender/pkg/api/render"
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/usecase/user"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	userUseCase user.IUseCase
}

func NewHandler(userUseCase user.IUseCase) *Handler {
	return &Handler{userUseCase: userUseCase}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var userDTO model.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		render.ResponseError(w, err, GenerateHTTPErrorStatusCode(errors.ErrInvalidPayload))
		return
	}

	u, err := h.userUseCase.Create(userDTO)
	if err != nil {
		render.ResponseError(w, err, GenerateHTTPErrorStatusCode(err))
		return
	}

	render.Response(w, u, http.StatusOK)
}

func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	u, err := h.userUseCase.Find()
	if err != nil {
		render.ResponseError(w, err, GenerateHTTPErrorStatusCode(err))
		return
	}

	render.Response(w, u, http.StatusOK)
}

func (h *Handler) FindById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	u, err := h.userUseCase.FindById(id)
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