package user

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/stretchr/testify/mock"
)

type UseCaseMock struct {
	mock.Mock
}

func (u *UseCaseMock) Create(userDto model.UserDto) (*model.User, error) {
	args := u.Called(userDto)
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (u *UseCaseMock) Find() (*model.Users, error) {
	args := u.Called()
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.Users), args.Error(1)
	}

	return nil, args.Error(1)
}

func (u *UseCaseMock) FindById(id string) (*model.User, error) {
	args := u.Called(id)
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.User), args.Error(1)
	}

	return nil, args.Error(1)
}
