package book

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Persist(book model.Book) (*model.Book, error) {
	args := r.Called(book)
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.Book), args.Error(1)
	}

	return nil, args.Error(1)
}
