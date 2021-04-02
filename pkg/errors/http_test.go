package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const dummyMsg = "dummy error"

var dummyErr = errors.New(dummyMsg)

func TestNotFoundError(t *testing.T) {
	err := NotFound{dummyErr}
	msg := err.Error()
	assert.Equal(t, dummyMsg, msg)
	assert.Exactly(t, dummyErr, err.Err)
}

func TestInvalidPayloadError(t *testing.T) {
	err := InvalidPayload{dummyErr}
	msg := err.Error()
	assert.Equal(t, dummyMsg, msg)
	assert.Exactly(t, dummyErr, err.Err)
}

func TestInternalServerError(t *testing.T) {
	err := InternalServerError{dummyErr}
	msg := err.Error()
	assert.Equal(t, dummyMsg, msg)
	assert.Exactly(t, dummyErr, err.Err)
}

func TestConflictError(t *testing.T) {
	err := Conflict{dummyErr}
	msg := err.Error()
	assert.Equal(t, dummyMsg, msg)
	assert.Exactly(t, dummyErr, err.Err)
}

func TestUnauthorizedError(t *testing.T) {
	err := Unauthorized{dummyErr}
	msg := err.Error()
	assert.Equal(t, dummyMsg, msg)
	assert.Exactly(t, dummyErr, err.Err)
}

func TestBadRequestError(t *testing.T) {
	err := BadRequest{dummyErr}
	msg := err.Error()
	assert.Equal(t, dummyMsg, msg)
	assert.Exactly(t, dummyErr, err.Err)
}
