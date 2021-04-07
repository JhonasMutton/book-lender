package user

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Persist(user model.User) (*model.User, error) {
	args := r.Called(user)
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *RepositoryMock) Fetch() (*model.Users, error) {
	args := r.Called()
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.Users), args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *RepositoryMock) FetchById(id uint) (*model.User, error) {
	args := r.Called(id)
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.User), args.Error(1)
	}

	return nil, args.Error(1)
}
