package lend

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Persist(loanBook model.LoanBook) (*model.LoanBook, error) {
	args := r.Called(loanBook)
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.LoanBook), args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *RepositoryMock) Update(loanBook model.LoanBook) (*model.LoanBook, error) {
	args := r.Called(loanBook)
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.LoanBook), args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *RepositoryMock) FetchByToUserAndBookAndStatus(toUserId, bookId uint, status string) (*model.LoanBook, error) {
	args := r.Called(toUserId, bookId, status)
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.LoanBook), args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *RepositoryMock) FetchByBookAndStatus(bookId uint, status string) (*model.LoanBook, error) {
	args := r.Called(bookId, status)
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.LoanBook), args.Error(1)
	}

	return nil, args.Error(1)
}
