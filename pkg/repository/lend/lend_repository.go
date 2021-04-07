package lend

import (
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/log"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"gorm.io/gorm"
)

type IRepository interface {
	Persist(loanBook model.LoanBook) (*model.LoanBook, error)
	Update(loanBook model.LoanBook) (*model.LoanBook, error)
	FetchByToUserAndBookAndStatus(toUserId, bookId uint, status string) (*model.LoanBook, error)
	FetchByBookAndStatus(bookId uint, status string) (*model.LoanBook, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Persist(loanBook model.LoanBook) (*model.LoanBook, error) {
	log.Logger.Debugf("Persisting book:%x, to user:%x", loanBook.Book, loanBook.ToUser)
	result := r.db.Create(&loanBook)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &loanBook, nil
}

func (r Repository) Update(loanBook model.LoanBook) (*model.LoanBook, error) {
	log.Logger.Debugf("Updating loan book:%x", loanBook.ID)
	result := r.db.Updates(&loanBook)
	if err := result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return &loanBook, nil
}

func (r Repository) FetchByToUserAndBookAndStatus(toUserId, bookId uint, status string) (*model.LoanBook, error) {
	log.Logger.Debugf("Fetching by toUser by book and status:%x, %x, %x", toUserId, bookId, status)
	var loanBook model.LoanBook
	result := r.db.Where("book = ? and to_user = ? and status = ?", bookId, toUserId, status).First(&loanBook)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &loanBook, nil
}

func (r Repository) FetchByBookAndStatus(bookId uint, status string) (*model.LoanBook, error) {
	log.Logger.Debugf("Fetching by book and status:%x, %x", bookId, status)
	var loanBook model.LoanBook
	result := r.db.Where("book = ? and status = ?", bookId, status).First(&loanBook)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &loanBook, nil
}
