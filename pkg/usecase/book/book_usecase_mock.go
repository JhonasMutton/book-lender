package book

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/stretchr/testify/mock"
)

type UseCaseMock struct {
	mock.Mock
}

func (u *UseCaseMock) Create(bookDto model.BookDTO) (*model.Book, error) {
	args := u.Called(bookDto)
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.Book), args.Error(1)
	}

	return nil, args.Error(1)
}
