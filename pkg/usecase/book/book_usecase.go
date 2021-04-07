package book

import (
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/log"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/book"
	"github.com/JhonasMutton/book-lender/pkg/validate"
)

type IUseCase interface {
	Create(bookDto model.BookDTO) (*model.Book, error)
}

type UseCase struct {
	bookRepository book.IRepository
	validator      *validate.Validator
}

func NewUseCase(bookRepository book.IRepository, validator *validate.Validator) *UseCase {
	return &UseCase{bookRepository: bookRepository, validator: validator}
}

func (u UseCase) Create(bookDto model.BookDTO) (*model.Book, error) {
	log.Logger.Debugf("Creating book: %s", bookDto.Title)
	if err := u.validator.Validate(bookDto); err != nil {
		return nil, errors.WrapWithMessage(errors.ErrInvalidPayload, err.Error())
	}

	bookModel := bookDto.ToModel()

	persisted, err := u.bookRepository.Persist(bookModel)
	if err != nil {
		return nil, errors.BuildError(err)
	}

	return persisted, nil
}
