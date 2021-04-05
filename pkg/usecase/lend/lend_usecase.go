package lend

import (
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/lend"
	"github.com/go-playground/validator"
	goErrors "github.com/pkg/errors"
	"gorm.io/gorm"

	"time"
)

type IUseCase interface {
	Lend(lendDTO model.LendBookDTO) (*model.LoanBook, error)
	Return(returnDTO model.ReturnBookDTO) (*model.LoanBook, error)
}

type UseCase struct {
	lendRepository *lend.Repository
	validate       *validator.Validate
}

func NewUseCase(lendRepository *lend.Repository, validate *validator.Validate) *UseCase {
	return &UseCase{lendRepository: lendRepository, validate: validate}
}

func (u UseCase) Lend(lendDTO model.LendBookDTO) (*model.LoanBook, error) {
	if err := u.validate.Struct(lendDTO); err != nil {
		return nil, err
	}

	lendModel := lendDTO.ToModel()

	loanBookFound, err := u.lendRepository.FetchByBookAndStatus(lendModel.Book, model.StatusLent)
	if err != nil && !goErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if loanBookFound != nil { //REGRA 3
		return nil, errors.New("book already lent")
	}

	lendModel.LentAt = time.Now()

	persisted, err := u.lendRepository.Persist(lendModel)
	if err != nil {
		return nil, err //TODO Handle errors
	}

	return persisted, nil
}

func (u UseCase) Return(returnDTO model.ReturnBookDTO) (*model.LoanBook, error) {
	if err := u.validate.Struct(returnDTO); err != nil {
		return nil, err
	}

	returnModel := returnDTO.ToModel()

	loanBookFound, err := u.lendRepository.FetchByToUserAndBookAndStatus(returnModel.ToUser, returnModel.Book, model.StatusLent)
	if err != nil {
		return nil, err
	}
	if loanBookFound == nil { //REGRA 4
		return nil, errors.New("book already returned")
	}

	loanBookFound.ReturnedAt = time.Now()
	loanBookFound.Status = model.StatusReturned

	persisted, err := u.lendRepository.Update(*loanBookFound)
	if err != nil {
		return nil, err //TODO Handle errors
	}

	return persisted, nil
}
