package lend

import (
	"bytes"
	"encoding/json"
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/usecase/lend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	lendMethodName   = "Lend"
	returnMethodName = "Return"
)

var errorScenarios = []struct {
	name           string
	err            error
	expectedStatus int
}{
	{"InvalidPayload", errors.ErrInvalidPayload, http.StatusPreconditionFailed},
	{"NotFound", errors.ErrNotFound, http.StatusNotFound},
	{"BadRequest", errors.ErrBadRequest, http.StatusBadRequest},
	{"InternalServerError", errors.ErrInternalServer, http.StatusInternalServerError},
	{"NonMapped", errors.New("dummy error"), http.StatusInternalServerError},
}

func TestNewHandler(t *testing.T) {
	useCaseMock := new(lend.UseCaseMock)

	handler := NewHandler(useCaseMock)

	assert.NotNil(t, handler)
	assert.Equal(t, useCaseMock, handler.lendUseCase)
}

func TestHandler_Lend(t *testing.T) {
	lendDto := model.LendBookDTO{
		Book:       22,
		LoggedUser: 11,
		ToUser:     33,
	}

	lendModel := lendDto.ToModel()
	lendModel.ID = 525

	useCaseMock := new(lend.UseCaseMock)
	useCaseMock.On(lendMethodName, mock.Anything).Return(&lendModel, nil)
	handler := NewHandler(useCaseMock)

	marshal, _ := json.Marshal(lendDto)
	req, _ := http.NewRequest(http.MethodPost, "/book/lend", bytes.NewBuffer(marshal))
	recorder := httptest.NewRecorder()

	handler.Lend(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "525")
}

func TestHandler_Lend_withInvalidBody(t *testing.T) {
	useCaseMock := new(lend.UseCaseMock)
	handler := NewHandler(useCaseMock)

	req, _ := http.NewRequest(http.MethodPost, "/book/lend", bytes.NewBuffer([]byte("test")))
	recorder := httptest.NewRecorder()

	handler.Lend(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandler_Lend_Errors(t *testing.T) {
	for _, s := range errorScenarios {
		t.Run(s.name, func(t *testing.T) {
			lendDto := model.LendBookDTO{
				Book:       22,
				LoggedUser: 11,
				ToUser:     33,
			}

			lendModel := lendDto.ToModel()
			lendModel.ID = 525

			useCaseMock := new(lend.UseCaseMock)
			useCaseMock.On(lendMethodName, mock.Anything).Return(nil, s.err)
			handler := NewHandler(useCaseMock)

			marshal, _ := json.Marshal(lendDto)
			req, _ := http.NewRequest(http.MethodPost, "/book/lend", bytes.NewBuffer(marshal))
			recorder := httptest.NewRecorder()

			handler.Lend(recorder, req)

			assert.Equal(t, s.expectedStatus, recorder.Code)
		})
	}
}

func TestHandler_Return(t *testing.T) {
	lendDto := model.ReturnBookDTO{
		Book:       22,
		LoggedUser: 11,
	}

	lendModel := lendDto.ToModel()
	lendModel.ID = 525

	useCaseMock := new(lend.UseCaseMock)
	useCaseMock.On(returnMethodName, mock.Anything).Return(&lendModel, nil)
	handler := NewHandler(useCaseMock)

	marshal, _ := json.Marshal(lendDto)
	req, _ := http.NewRequest(http.MethodPost, "/book/return", bytes.NewBuffer(marshal))
	recorder := httptest.NewRecorder()

	handler.Return(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "525")
}

func TestHandler_Return_withInvalidBody(t *testing.T) {
	useCaseMock := new(lend.UseCaseMock)
	handler := NewHandler(useCaseMock)

	req, _ := http.NewRequest(http.MethodPost, "/book/return", bytes.NewBuffer([]byte("test")))
	recorder := httptest.NewRecorder()

	handler.Return(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandler_Return_Errors(t *testing.T) {
	for _, s := range errorScenarios {
		t.Run(s.name, func(t *testing.T) {
			lendDto := model.LendBookDTO{
				Book:       22,
				LoggedUser: 11,
				ToUser:     33,
			}

			lendModel := lendDto.ToModel()
			lendModel.ID = 525

			useCaseMock := new(lend.UseCaseMock)
			useCaseMock.On(returnMethodName, mock.Anything).Return(nil, s.err)
			handler := NewHandler(useCaseMock)

			marshal, _ := json.Marshal(lendDto)
			req, _ := http.NewRequest(http.MethodPost, "/book/lend", bytes.NewBuffer(marshal))
			recorder := httptest.NewRecorder()

			handler.Return(recorder, req)

			assert.Equal(t, s.expectedStatus, recorder.Code)
		})
	}
}
