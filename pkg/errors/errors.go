package errors

import (
	"fmt"
	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var ErrNotFound = &NotFound{Err: fmt.Errorf("not found")}
var ErrUnauthorized = &Unauthorized{Err: fmt.Errorf("unauthorized")}
var ErrInvalidPayload = &InvalidPayload{Err: fmt.Errorf("invalid payload")}
var ErrInternalServer = &InternalServerError{Err: fmt.Errorf("internal server error")}
var ErrBadRequest = &BadRequest{Err: fmt.Errorf("internal server error")}
var ErrConflict = &Conflict{Err: fmt.Errorf("conflict")}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func Wrap(err error) error {
	if _, ok := err.(stackTracer); ok {
		return err
	}

	return errors.WithStack(err)
}

func WrapWithMessage(err error, msg string) error {
	if _, ok := err.(stackTracer); ok {
		return errors.WithMessage(err, msg)
	}

	return errors.Wrap(err, msg)
}

func Cause(err error) error {
	return errors.Cause(err)
}

func New(msg string) error {
	return errors.New(msg)
}

func BuildError(err error) error {
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		switch driverErr.Number {
		case mysqlerr.ER_DUP_ENTRY:
			return WrapWithMessage(ErrConflict, "some key already in use")
		}
	}

	if errors.Is(err, gorm.ErrRecordNotFound){
		return WrapWithMessage(ErrNotFound, err.Error())
	}

	return WrapWithMessage(ErrInternalServer, "some error has occurred")
}
