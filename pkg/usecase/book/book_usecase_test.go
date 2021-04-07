package book

import (
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/book"
	"github.com/JhonasMutton/book-lender/pkg/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	v = validate.NewValidator()
)

const (
	PersistMethodName = "Persist"
)

func TestNewUseCase(t *testing.T) {
	repoMock := new(book.RepositoryMock)

	useCase := NewUseCase(repoMock, v)
	assert.NotNil(t, useCase)
	assert.Equal(t, repoMock, useCase.bookRepository)
	assert.Equal(t, v, useCase.validator)
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
	useCase := NewUseCase(repoMock, v)

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
	useCase := NewUseCase(repoMock, v)

	persisted, err := useCase.Create(bookDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "Pages is a required field: invalid payload")
}

func TestUseCase_Create_withPersistError(t *testing.T) {
	bookDto := model.BookDTO{
		Title:        "Sherlock Holmes - Um estudo em vermelho",
		Pages:        "176",
		LoggedUserId: 10,
	}

	repoMock := new(book.RepositoryMock)
	repoMock.On(PersistMethodName, mock.Anything).Return(nil, errors.New("error to persist"))
	useCase := NewUseCase(repoMock, v)

	persisted, err := useCase.Create(bookDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "some error has occurred: internal server error")
}
