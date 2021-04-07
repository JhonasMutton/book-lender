package lend

import (
	"encoding/json"
	"github.com/JhonasMutton/book-lender/pkg/api/render"
	"github.com/JhonasMutton/book-lender/pkg/log"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/usecase/lend"
	"net/http"
)

type Handler struct {
	lendUseCase lend.IUseCase
}

func NewHandler(lendUseCase lend.IUseCase) *Handler {
	return &Handler{lendUseCase: lendUseCase}
}

func (h *Handler) Lend(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("Handling lend book!")
	var lendDTO model.LendBookDTO
	if err := json.NewDecoder(r.Body).Decode(&lendDTO); err != nil {
		log.Logger.Errorf("Error to decode entry, error: %s", err.Error())
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	u, err := h.lendUseCase.Lend(lendDTO)
	if err != nil {
		log.Logger.Errorf("Some error has occurred on Lend, error: %s", err.Error())
		render.ResponseError(w, err, render.GenerateHTTPErrorStatusCode(err))
		return
	}

	render.Response(w, u, http.StatusOK)
	log.Logger.Info("Lend book was handled")
}

func (h *Handler) Return(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("Handling return book!")
	var returnDTO model.ReturnBookDTO
	if err := json.NewDecoder(r.Body).Decode(&returnDTO); err != nil {
		log.Logger.Errorf("Error to decode entry, error: %s", err.Error())
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	u, err := h.lendUseCase.Return(returnDTO)
	if err != nil {
		log.Logger.Errorf("Some error has occurred on Return, error: %s", err.Error())
		render.ResponseError(w, err, render.GenerateHTTPErrorStatusCode(err))
		return
	}

	render.Response(w, u, http.StatusOK)
	log.Logger.Info("Return book was handled")
}
