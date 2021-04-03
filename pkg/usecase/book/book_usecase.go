package book

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/book"
)

type IUseCase interface {
	Create(bookDto model.BookDTO) (*model.Book, error)
}

type UseCase struct {
	bookRepository book.IRepository
}

func NewUseCase(bookRepository book.IRepository) *UseCase {
	return &UseCase{bookRepository: bookRepository}
}

func (u UseCase) Create(bookDto model.BookDTO) (*model.Book, error) {
	//TODO VALIDATE FIELDS
	bookModel := bookDto.ToModel()

	persisted, err := u.bookRepository.Persist(bookModel)
	if err != nil {
		return nil, err //TODO Handle errors
	}

	return persisted, nil
}
