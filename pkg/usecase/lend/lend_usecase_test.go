package lend

import (
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/lend"
	 "github.com/JhonasMutton/book-lender/pkg/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	v = validate.NewValidator()
)

const (
	PersistMethodName       = "Persist"
	FetchByToUserMethodName = "FetchByToUserAndBookAndStatus"
	FetchByBookMethodName   = "FetchByBookAndStatus"
	UpdateMethodName        = "Update"
)

func TestNewUseCase(t *testing.T) {
	repoMock := new(lend.RepositoryMock)

	useCase := NewUseCase(repoMock, v)
	assert.NotNil(t, useCase)
	assert.Equal(t, repoMock, useCase.lendRepository)
	assert.Equal(t, v, useCase.validator)
}

func TestUseCase_Lend(t *testing.T) {
	lendDto := model.LendBookDTO{
		Book:       22,
		LoggedUser: 11,
		ToUser:     33,
	}

	lendModel := lendDto.ToModel()
	lendModel.ID = 10

	repoMock := new(lend.RepositoryMock)
	repoMock.On(FetchByBookMethodName, mock.Anything, mock.Anything).Return(nil, nil)
	repoMock.On(PersistMethodName, mock.Anything).Return(&lendModel, nil)
	useCase := NewUseCase(repoMock, v)

	persisted, err := useCase.Lend(lendDto)

	assert.NoError(t, err)
	assert.NotNil(t, persisted)
	assert.NotZero(t, persisted.ID)
}

func TestUseCase_Lend_withValidateError(t *testing.T) {
	lendDto := model.LendBookDTO{
		Book:       22,
		LoggedUser: 11,
	}

	repoMock := new(lend.RepositoryMock)
	useCase := NewUseCase(repoMock, v)

	persisted, err := useCase.Lend(lendDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "ToUser is a required field: invalid payload")
}

func TestUseCase_Lend_withPersistError(t *testing.T) {
	lendDto := model.LendBookDTO{
		Book:       22,
		LoggedUser: 11,
		ToUser:     33,
	}

	repoMock := new(lend.RepositoryMock)
	repoMock.On(FetchByBookMethodName, mock.Anything, mock.Anything).Return(nil, nil)
	repoMock.On(PersistMethodName, mock.Anything).Return(nil, errors.New("error to persist"))
	useCase := NewUseCase(repoMock, v)

	persisted, err := useCase.Lend(lendDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "some error has occurred: internal server error")
}

func TestUseCase_Return(t *testing.T) {
	returnDto := model.ReturnBookDTO{
		Book:       10,
		LoggedUser: 11,
	}

	returnModel := returnDto.ToModel()
	returnModel.ID = 10

	repoMock := new(lend.RepositoryMock)
	repoMock.On(FetchByToUserMethodName, mock.Anything, mock.Anything, mock.Anything).Return(&returnModel, nil)
	repoMock.On(UpdateMethodName, mock.Anything).Return(&returnModel, nil)
	useCase := NewUseCase(repoMock, v)

	persisted, err := useCase.Return(returnDto)

	assert.NoError(t, err)
	assert.NotNil(t, persisted)
	assert.NotZero(t, persisted.ID)
}

func TestUseCase_Return_withValidateError(t *testing.T) {
	returnDto := model.ReturnBookDTO{
		Book: 22,
	}

	repoMock := new(lend.RepositoryMock)
	useCase := NewUseCase(repoMock, v)

	persisted, err := useCase.Return(returnDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "LoggedUser is a required field")
}

func TestUseCase_Return_withUpdateError(t *testing.T) {
	returnDto := model.ReturnBookDTO{
		Book:       22,
		LoggedUser: 11,
	}

	returnModel := returnDto.ToModel()

	repoMock := new(lend.RepositoryMock)
	repoMock.On(FetchByToUserMethodName, mock.Anything, mock.Anything, mock.Anything).Return(&returnModel, nil)
	repoMock.On(UpdateMethodName, mock.Anything).Return(nil, errors.New("error to update"))
	useCase := NewUseCase(repoMock, v)

	persisted, err := useCase.Return(returnDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "some error has occurred: internal server error")
}
