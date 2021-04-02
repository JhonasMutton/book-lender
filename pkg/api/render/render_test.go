package render

import (
	"bitbucket.org/sensedia/secret-manager/pkg/test/testdata/utils"
	"errors"
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
	event := utils.ReadFile("../../test/testdata/json/payload/secret/http_response_status.json")

	// when
	Response(recorder, event, http.StatusOK)

	// then
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusOK, recorder.Code)
}
