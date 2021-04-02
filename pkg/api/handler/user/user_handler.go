package user

import (
	"encoding/json"
	"github.com/JhonasMutton/book-lender/pkg/api/render"
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/usecase/user"
	"net/http"
)

type Handler struct {
	userUseCase user.IUseCase
}

func NewHandler(userUseCase user.IUseCase) *Handler {
	return &Handler{userUseCase: userUseCase}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var basicUser model.BasicUser
	if err := json.NewDecoder(r.Body).Decode(&basicUser); err != nil {
		render.ResponseError(w, err, GenerateHTTPErrorStatusCode(errors.ErrInvalidPayload))
	}

	u, err := h.userUseCase.CreateUser(basicUser)
	if err != nil {
		render.ResponseError(w, err, GenerateHTTPErrorStatusCode(err))
		return
	}

	render.Response(w, u, http.StatusOK)
}

func (h *Handler) FindUsers(w http.ResponseWriter, r *http.Request) {
	u, err := h.userUseCase.FindUsers()
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