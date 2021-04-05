package book

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/book"
	"github.com/go-playground/validator"
)

type IUseCase interface {
	Create(bookDto model.BookDTO) (*model.Book, error)
}

type UseCase struct {
	bookRepository book.IRepository
	validate       *validator.Validate
}

func NewUseCase(bookRepository book.IRepository, validate *validator.Validate) *UseCase {
	return &UseCase{bookRepository: bookRepository, validate: validate}
}

func (u UseCase) Create(bookDto model.BookDTO) (*model.Book, error) {
	if err := u.validate.Struct(bookDto); err != nil {
		return nil, err
	}

	bookModel := bookDto.ToModel()

	persisted, err := u.bookRepository.Persist(bookModel)
	if err != nil {
		return nil, err //TODO Handle errors
	}

	return persisted, nil
}
