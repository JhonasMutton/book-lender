package render

import (
	"encoding/json"
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, response interface{}, code int) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		ResponseError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(jsonResponse)
}

func ResponseError(w http.ResponseWriter, err error, code int) {
	jsonResponse, err := json.Marshal(ErrorResponse{code, err.Error()})
	if err != nil {
		ResponseError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(jsonResponse)
}

func GenerateHTTPErrorStatusCode(err error) int {
	switch errors.Cause(err).(type) {
	case *errors.NotFound:
		return http.StatusNotFound
	case *errors.InvalidPayload:
		return http.StatusPreconditionFailed
	case *errors.BadRequest:
		return http.StatusBadRequest
	case *errors.Conflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}