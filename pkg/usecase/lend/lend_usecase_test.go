package lend

import (
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/lend"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	validate = validator.New()
)

const (
	PersistMethodName       = "Persist"
	FetchByToUserMethodName = "FetchByToUserAndBookAndStatus"
	FetchByBookMethodName   = "FetchByBookAndStatus"
	UpdateMethodName = "Update"
)

func TestNewUseCase(t *testing.T) {
	repoMock := new(lend.RepositoryMock)

	useCase := NewUseCase(repoMock, validate)
	assert.NotNil(t, useCase)
	assert.Equal(t, repoMock, useCase.lendRepository)
	assert.Equal(t, validate, useCase.validate)
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
	useCase := NewUseCase(repoMock, validate)

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
	useCase := NewUseCase(repoMock, validate)

	persisted, err := useCase.Lend(lendDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "Key: 'LendBookDTO.ToUser' Error:Field validation for 'ToUser' failed on the 'required' tag")
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
	useCase := NewUseCase(repoMock, validate)

	persisted, err := useCase.Lend(lendDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "error to persist")
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
	useCase := NewUseCase(repoMock, validate)

	persisted, err := useCase.Return(returnDto)

	assert.NoError(t, err)
	assert.NotNil(t, persisted)
	assert.NotZero(t, persisted.ID)
}

func TestUseCase_Return_withValidateError(t *testing.T) {
	returnDto := model.ReturnBookDTO{
		Book:       22,
	}

	repoMock := new(lend.RepositoryMock)
	useCase := NewUseCase(repoMock, validate)

	persisted, err := useCase.Return(returnDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "Key: 'ReturnBookDTO.LoggedUser' Error:Field validation for 'LoggedUser' failed on the 'required' tag")
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
	useCase := NewUseCase(repoMock, validate)

	persisted, err := useCase.Return(returnDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "error to update")
}
