package book

import (
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/book"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	validate = validator.New()
)

const (
	PersistMethodName = "Persist"
)

func TestNewUseCase(t *testing.T) {
	repoMock := new(book.RepositoryMock)

	useCase := NewUseCase(repoMock, validate)
	assert.NotNil(t, useCase)
	assert.Equal(t, repoMock, useCase.bookRepository)
	assert.Equal(t, validate, useCase.validate)
}

func TestUseCase_Create(t *testing.T) {
	bookDto := model.BookDTO{
		Title:        "Sherlock Holmes - Um estudo em vermelho",
		Pages:        "176",
		LoggedUserId: 10,
	}

	bookModel := bookDto.ToModel()
	bookModel.ID = 10

	repoMock := new(book.RepositoryMock)
	repoMock.On(PersistMethodName, mock.Anything).Return(&bookModel, nil)
	useCase := NewUseCase(repoMock, validate)
	
	persisted, err := useCase.Create(bookDto)

	assert.NoError(t, err)
	assert.NotNil(t, persisted)
	assert.NotZero(t, persisted.ID)
}

func TestUseCase_Create_withValidateError(t *testing.T) {
	bookDto := model.BookDTO{
		Title:        "Sherlock Holmes - Um estudo em vermelho",
		Pages:        "",
		LoggedUserId: 10,
	}

	repoMock := new(book.RepositoryMock)
	useCase := NewUseCase(repoMock, validate)

	persisted, err := useCase.Create(bookDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "Key: 'BookDTO.Pages' Error:Field validation for 'Pages' failed on the 'required' tag")
}


func TestUseCase_Create_withPersistError(t *testing.T) {
	bookDto := model.BookDTO{
		Title:        "Sherlock Holmes - Um estudo em vermelho",
		Pages:        "176",
		LoggedUserId: 10,
	}

	repoMock := new(book.RepositoryMock)
	repoMock.On(PersistMethodName, mock.Anything).Return(nil, errors.New("error to persist"))
	useCase := NewUseCase(repoMock, validate)

	persisted, err := useCase.Create(bookDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "error to persist")
}
