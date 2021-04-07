package lend

import (
	"encoding/json"
	"github.com/JhonasMutton/book-lender/pkg/api/render"
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
	var lendDTO model.LendBookDTO
	if err := json.NewDecoder(r.Body).Decode(&lendDTO); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	u, err := h.lendUseCase.Lend(lendDTO)
	if err != nil {
		render.ResponseError(w, err, render.GenerateHTTPErrorStatusCode(err))
		return
	}

	render.Response(w, u, http.StatusOK)
}

func (h *Handler) Return(w http.ResponseWriter, r *http.Request) {
	var returnDTO model.ReturnBookDTO
	if err := json.NewDecoder(r.Body).Decode(&returnDTO); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	u, err := h.lendUseCase.Return(returnDTO)
	if err != nil {
		render.ResponseError(w, err, render.GenerateHTTPErrorStatusCode(err))
		return
	}

	render.Response(w, u, http.StatusOK)
}
