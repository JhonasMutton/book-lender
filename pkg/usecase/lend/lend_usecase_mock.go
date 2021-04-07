package lend

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/stretchr/testify/mock"
)

type UseCaseMock struct {
	mock.Mock
}

func (u *UseCaseMock) Lend(lendDTO model.LendBookDTO) (*model.LoanBook, error) {
	args := u.Called(lendDTO)
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.LoanBook), args.Error(1)
	}

	return nil, args.Error(1)
}

func (u *UseCaseMock) Return(returnDTO model.ReturnBookDTO) (*model.LoanBook, error) {
	args := u.Called(returnDTO)
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.(*model.LoanBook), args.Error(1)
	}

	return nil, args.Error(1)
}
