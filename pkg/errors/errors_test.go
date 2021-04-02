package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrap(t *testing.T) {
	err := Wrap(dummyErr)
	assert.NotNil(t, err)
}

func TestWrapWithMessage(t *testing.T) {
	err := WrapWithMessage(dummyErr, dummyMsg)
	assert.NotNil(t, err)
}

func TestCause(t *testing.T) {
	err := Cause(dummyErr)
	assert.NotNil(t, err)
}

func TestNew(t *testing.T) {
	err := New(dummyMsg)
	assert.NotNil(t, err)
	assert.Equal(t, dummyMsg, err.Error())
}
