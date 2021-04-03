package lend

import (
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/lend"
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
}

func NewUseCase(lendRepository *lend.Repository) *UseCase {
	return &UseCase{lendRepository: lendRepository}
}

func (u UseCase) Lend(lendDTO model.LendBookDTO) (*model.LoanBook, error) {
	//TODO VALIDATE FIELDS
	lendModel := lendDTO.ToModel()

	loanBookFound, err := u.lendRepository.FindByUsers(lendModel)
	if err != nil && !goErrors.Is(err, gorm.ErrRecordNotFound){
		return nil, err
	}
	if loanBookFound != nil {
		return nil, errors.New("book already lent")
	}

	lendModel.LentAt =  time.Now()

	persisted, err := u.lendRepository.Persist(lendModel)
	if err != nil {
		return nil, err //TODO Handle errors
	}

	return persisted, nil
}

func (u UseCase) Return(returnDTO model.ReturnBookDTO) (*model.LoanBook, error) {
	//TODO VALIDATE FIELDS
	returnModel := returnDTO.ToModel()

	loanBookFound, err := u.lendRepository.FindByToUser(returnModel)
	if err != nil {
		return nil, err
	}
	if loanBookFound == nil {
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
