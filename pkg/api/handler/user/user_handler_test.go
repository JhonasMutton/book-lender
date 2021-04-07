package user

import (
	"bytes"
	"encoding/json"
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/log"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/usecase/user"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	log.SetupLogger()
}

const (
	createMethodName   = "Create"
	findMethodName     = "Find"
	findByIdMethodName = "FindById"
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
	useCaseMock := new(user.UseCaseMock)

	handler := NewHandler(useCaseMock)

	assert.NotNil(t, handler)
	assert.Equal(t, useCaseMock, handler.userUseCase)
}

func TestHandler_Post(t *testing.T) {
	userDto := model.UserDto{
		Name:  "Alfredão",
		Email: "ofred.z1k4@hotmail.com",
	}

	userModel := userDto.ToModel()
	userModel.ID = 525

	useCaseMock := new(user.UseCaseMock)
	useCaseMock.On(createMethodName, mock.Anything).Return(&userModel, nil)
	handler := NewHandler(useCaseMock)

	marshal, _ := json.Marshal(userDto)
	req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(marshal))
	recorder := httptest.NewRecorder()

	handler.Post(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "525")
}

func TestHandler_Post_withInvalidBody(t *testing.T) {
	useCaseMock := new(user.UseCaseMock)
	handler := NewHandler(useCaseMock)

	req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer([]byte("test")))
	recorder := httptest.NewRecorder()

	handler.Post(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandler_Post_Errors(t *testing.T) {
	for _, s := range errorScenarios {
		t.Run(s.name, func(t *testing.T) {
			userDto := model.UserDto{
				Name:  "Batmão",
				Email: "batima.fernardes@hotmail.com",
			}

			userModel := userDto.ToModel()
			userModel.ID = 525

			useCaseMock := new(user.UseCaseMock)
			useCaseMock.On(createMethodName, mock.Anything).Return(nil, s.err)
			handler := NewHandler(useCaseMock)

			marshal, _ := json.Marshal(userDto)
			req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(marshal))
			recorder := httptest.NewRecorder()

			handler.Post(recorder, req)

			assert.Equal(t, s.expectedStatus, recorder.Code)
		})
	}
}

func TestHandler_Find(t *testing.T) {
	userDto := model.UserDto{
		Name:  "Alfredão",
		Email: "ofred.z1k4@hotmail.com",
	}

	userModel := userDto.ToModel()
	userModel.ID = 525
	users := model.Users{userModel}

	useCaseMock := new(user.UseCaseMock)
	useCaseMock.On(findMethodName).Return(&users, nil)
	handler := NewHandler(useCaseMock)

	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	recorder := httptest.NewRecorder()

	handler.Find(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "525")
}

func TestHandler_Find_Errors(t *testing.T) {
	for _, s := range errorScenarios {
		t.Run(s.name, func(t *testing.T) {
			useCaseMock := new(user.UseCaseMock)
			useCaseMock.On(findMethodName).Return(nil, s.err)
			handler := NewHandler(useCaseMock)

			req, _ := http.NewRequest(http.MethodPost, "/user", nil)
			recorder := httptest.NewRecorder()

			handler.Find(recorder, req)

			assert.Equal(t, s.expectedStatus, recorder.Code)
		})
	}
}

func TestHandler_FindById(t *testing.T) {
	userDto := model.UserDto{
		Name:  "Alfredão",
		Email: "ofred.z1k4@hotmail.com",
	}

	userModel := userDto.ToModel()
	userModel.ID = 525

	useCaseMock := new(user.UseCaseMock)
	useCaseMock.On(findByIdMethodName, mock.Anything).Return(&userModel, nil)
	handler := NewHandler(useCaseMock)

	req, _ := http.NewRequest(http.MethodGet, "/user/{id}", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "525",
	})

	recorder := httptest.NewRecorder()

	handler.FindById(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "525")
}

func TestHandler_FindById_Errors(t *testing.T) {
	for _, s := range errorScenarios {
		t.Run(s.name, func(t *testing.T) {
			useCaseMock := new(user.UseCaseMock)
			useCaseMock.On(findByIdMethodName, mock.Anything).Return(nil, s.err)
			handler := NewHandler(useCaseMock)

			req, _ := http.NewRequest(http.MethodPost, "/user/{id}", nil)
			recorder := httptest.NewRecorder()

			handler.FindById(recorder, req)

			assert.Equal(t, s.expectedStatus, recorder.Code)
		})
	}
}
