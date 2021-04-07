package user

import (
	"encoding/json"
	"github.com/JhonasMutton/book-lender/pkg/api/render"
	"github.com/JhonasMutton/book-lender/pkg/log"
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
	log.Logger.Info("Handling user post!")
	var userDTO model.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		log.Logger.Errorf("Error to decode entry, error: %s", err.Error())
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	u, err := h.userUseCase.Create(userDTO)
	if err != nil {
		log.Logger.Errorf("Some error has occurred on Create, error: %s", err.Error())
		render.ResponseError(w, err, render.GenerateHTTPErrorStatusCode(err))
		return
	}

	render.Response(w, u, http.StatusCreated)
	log.Logger.Info("Post was handled")
}

func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("Handling find user!")
	u, err := h.userUseCase.Find()
	if err != nil {
		log.Logger.Errorf("Some error has occurred on Find, error: %s", err.Error())
		render.ResponseError(w, err, render.GenerateHTTPErrorStatusCode(err))
		return
	}

	render.Response(w, u, http.StatusOK)
	log.Logger.Info("Find was handled")
}

func (h *Handler) FindById(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("Handling find user by id!")
	id := mux.Vars(r)["id"]

	u, err := h.userUseCase.FindById(id)
	if err != nil {
		log.Logger.Errorf("Some error has occurred on FindById, error: %s", err.Error())
		render.ResponseError(w, err, render.GenerateHTTPErrorStatusCode(err))
		return
	}

	render.Response(w, u, http.StatusOK)
	log.Logger.Info("Find by id was handled")
}