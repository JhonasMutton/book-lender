package user

import (
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/user"
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
	FetchMethodName   = "Fetch"
	FetchByIdMethodName   = "FetchById"

)

func TestNewUseCase(t *testing.T) {
	repoMock := new(user.RepositoryMock)

	useCase := NewUseCase(repoMock, validate)
	assert.NotNil(t, useCase)
	assert.Equal(t, repoMock, useCase.userRepository)
	assert.Equal(t, validate, useCase.validate)
}

func TestUseCase_Create(t *testing.T) {
	userDto := model.UserDto{
		Name:  "Bruce Robertinho Wayne",
		Email: "bruce.wayne@wayne.com",
	}

	userModel := userDto.ToModel()
	userModel.ID = 10

	repoMock := new(user.RepositoryMock)
	repoMock.On(PersistMethodName, mock.Anything).Return(&userModel, nil)
	useCase := NewUseCase(repoMock, validate)

	persisted, err := useCase.Create(userDto)

	assert.NoError(t, err)
	assert.NotNil(t, persisted)
	assert.NotZero(t, persisted.ID)
}

func TestUseCase_Create_withValidateError(t *testing.T) {
	userDto := model.UserDto{
		Name:  "Bruce Robertinho Wayne",
		Email: "bruce.com",
	}

	repoMock := new(user.RepositoryMock)
	useCase := NewUseCase(repoMock, validate)

	persisted, err := useCase.Create(userDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "Key: 'UserDto.Email' Error:Field validation for 'Email' failed on the 'email' tag")
}

func TestUseCase_Create_withPersistError(t *testing.T) {
	userDto := model.UserDto{
		Name:  "Bruce Robertinho Wayne",
		Email: "bruce.wayne@wayne.com",
	}

	repoMock := new(user.RepositoryMock)
	repoMock.On(PersistMethodName, mock.Anything).Return(nil, errors.New("error to persist"))
	useCase := NewUseCase(repoMock, validate)

	persisted, err := useCase.Create(userDto)

	assert.Nil(t, persisted)
	assert.Error(t, err)
	assert.EqualError(t, err, "error to persist")
}

func TestUseCase_Find(t *testing.T) {
	userDto := model.UserDto{
		Name:  "Bruce Robertinho Wayne",
		Email: "bruce.wayne@wayne.com",
	}

	users := model.Users{userDto.ToModel()}
	repoMock := new(user.RepositoryMock)
	repoMock.On(FetchMethodName).Return(&users, nil)
	useCase := NewUseCase(repoMock, validate)

	found, err := useCase.Find()

	assert.NotNil(t, found)
	assert.NotEmpty(t, found)
	assert.NoError(t, err)
}

func TestUseCase_Find_withFetchError(t *testing.T) {
	repoMock := new(user.RepositoryMock)
	repoMock.On(FetchMethodName).Return(nil, errors.New("fetch error"))
	useCase := NewUseCase(repoMock, validate)

	found, err := useCase.Find()

	assert.Nil(t, found)
	assert.Error(t, err)
	assert.EqualError(t, err, "fetch error")
}

func TestUseCase_Find_withNotFound(t *testing.T) {
	users := make(model.Users, 0)
	repoMock := new(user.RepositoryMock)
	repoMock.On(FetchMethodName).Return(&users, nil)
	useCase := NewUseCase(repoMock, validate)

	found, err := useCase.Find()

	assert.NotNil(t, found)
	assert.Empty(t, found)
	assert.NoError(t, err)
}

func TestUseCase_FindById(t *testing.T) {
	userDto := model.UserDto{
		Name:  "Bruce Robertinho Wayne",
		Email: "bruce.wayne@wayne.com",
	}
	userModel := userDto.ToModel()

	repoMock := new(user.RepositoryMock)
	repoMock.On(FetchByIdMethodName, mock.Anything).Return(&userModel, nil)
	useCase := NewUseCase(repoMock, validate)

	found, err := useCase.FindById("22")

	assert.NotNil(t, found)
	assert.NoError(t, err)
	assert.Equal(t, userDto.Name, found.Name)
}

func TestUseCase_FindById_withFetchError(t *testing.T) {
	repoMock := new(user.RepositoryMock)
	repoMock.On(FetchByIdMethodName, mock.Anything).Return(nil, errors.New("fetch by id error"))
	useCase := NewUseCase(repoMock, validate)

	found, err := useCase.FindById("22")

	assert.Nil(t, found)
	assert.Error(t, err)
	assert.EqualError(t, err, "fetch by id error")
}

func TestUseCase_FindById_withInvalidEntry(t *testing.T) {
	repoMock := new(user.RepositoryMock)
	useCase := NewUseCase(repoMock, validate)

	found, err := useCase.FindById("Xablau")

	assert.Nil(t, found)
	assert.Error(t, err)
	assert.EqualError(t, err, "strconv.ParseUint: parsing \"Xablau\": invalid syntax")
}