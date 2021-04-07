package render

import (
	"errors"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponseError(t *testing.T) {
	// given
	recorder := httptest.NewRecorder()
	err := errors.New("Dummy Error")

	// when
	ResponseError(recorder, err, http.StatusNotFound)

	// then
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestResponse(t *testing.T) {
	// given
	recorder := httptest.NewRecorder()
	userDto := model.UserDto{
		Name:  "Alfredão",
		Email: "ofred.z1k4@hotmail.com",
	}

	// when
	Response(recorder, userDto, http.StatusOK)

	// then
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusOK, recorder.Code)
}
