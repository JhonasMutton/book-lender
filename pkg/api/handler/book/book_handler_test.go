package book

import (
	"bytes"
	"encoding/json"
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/log"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/usecase/book"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

const createMethodName = "Create"

func init() {
	log.SetupLogger()
}

func TestNewHandler(t *testing.T) {
	useCaseMock := new(book.UseCaseMock)

	handler := NewHandler(useCaseMock)

	assert.NotNil(t, handler)
	assert.Equal(t, useCaseMock, handler.bookUseCase)
}

func TestHandler_Post(t *testing.T) {
	bookDto := model.BookDTO{
		Title:        "Sherlock Holmes - Um estudo em vermelho",
		Pages:        "176",
		LoggedUserId: 10,
	}

	bookModel := bookDto.ToModel()
	bookModel.ID = 525

	useCaseMock := new(book.UseCaseMock)
	useCaseMock.On(createMethodName, mock.Anything).Return(&bookModel, nil)
	handler := NewHandler(useCaseMock)

	marshal, _ := json.Marshal(bookDto)
	req, _ := http.NewRequest(http.MethodPost, "/book", bytes.NewBuffer(marshal))
	recorder := httptest.NewRecorder()

	handler.Post(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "525")
}

func TestHandler_Post_withInvalidBody(t *testing.T) {
	useCaseMock := new(book.UseCaseMock)
	handler := NewHandler(useCaseMock)

	req, _ := http.NewRequest(http.MethodPost, "/book", bytes.NewBuffer([]byte("test")))
	recorder := httptest.NewRecorder()

	handler.Post(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandler_Errors(t *testing.T) {
	scenarios := []struct {
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

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			bookDto := model.BookDTO{
				Title:        "Sherlock Holmes - Um estudo em vermelho",
				Pages:        "176",
				LoggedUserId: 10,
			}

			bookModel := bookDto.ToModel()
			bookModel.ID = 525

			useCaseMock := new(book.UseCaseMock)
			useCaseMock.On(createMethodName, mock.Anything).Return(nil, s.err)
			handler := NewHandler(useCaseMock)

			marshal, _ := json.Marshal(bookDto)
			req, _ := http.NewRequest(http.MethodPost, "/book", bytes.NewBuffer(marshal))
			recorder := httptest.NewRecorder()

			handler.Post(recorder, req)

			assert.Equal(t, s.expectedStatus, recorder.Code)
		})
	}
}