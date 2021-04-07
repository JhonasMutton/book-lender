package lend

import (
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/log"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/lend"
	"github.com/JhonasMutton/book-lender/pkg/validate"
	goErrors "github.com/pkg/errors"
	"gorm.io/gorm"

	"time"
)

type IUseCase interface {
	Lend(lendDTO model.LendBookDTO) (*model.LoanBook, error)
	Return(returnDTO model.ReturnBookDTO) (*model.LoanBook, error)
}

type UseCase struct {
	lendRepository lend.IRepository
	validator      *validate.Validator
}

func NewUseCase(lendRepository lend.IRepository, validator *validate.Validator) *UseCase {
	return &UseCase{lendRepository: lendRepository, validator: validator}
}

func (u UseCase) Lend(lendDTO model.LendBookDTO) (*model.LoanBook, error) {
	log.Logger.Debugf("Lending book %x to user %x", lendDTO.Book, lendDTO.ToUser)
	if err := u.validator.Validate(lendDTO); err != nil {
		return nil, errors.WrapWithMessage(errors.ErrInvalidPayload, err.Error())
	}

	lendModel := lendDTO.ToModel()

	loanBookFound, err := u.lendRepository.FetchByBookAndStatus(lendModel.Book, model.StatusLent)
	if err != nil && !goErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.BuildError(err)
	}
	if loanBookFound != nil { //REGRA 3
		return nil, errors.WrapWithMessage(errors.ErrConflict, "book already lent")
	}

	lendModel.LentAt = time.Now()

	persisted, err := u.lendRepository.Persist(lendModel)
	if err != nil {
		return nil, errors.BuildError(err)
	}

	return persisted, nil
}

func (u UseCase) Return(returnDTO model.ReturnBookDTO) (*model.LoanBook, error) {
	log.Logger.Debugf("Returning book %x from user %x", returnDTO.Book, returnDTO.LoggedUser)
	if err := u.validator.Validate(returnDTO); err != nil {
		return nil, err
	}

	returnModel := returnDTO.ToModel()

	loanBookFound, err := u.lendRepository.FetchByToUserAndBookAndStatus(returnModel.ToUser, returnModel.Book, model.StatusLent)
	if err != nil && !goErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.BuildError(err)
	}

	if loanBookFound == nil || goErrors.Is(err, gorm.ErrRecordNotFound) { //REGRA 4
		return nil, errors.WrapWithMessage(errors.ErrConflict, "book already returned")
	}

	loanBookFound.ReturnedAt = time.Now()
	loanBookFound.Status = model.StatusReturned

	persisted, err := u.lendRepository.Update(*loanBookFound)
	if err != nil {
		return nil, errors.BuildError(err)
	}

	return persisted, nil
}
