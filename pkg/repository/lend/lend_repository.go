package lend

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"gorm.io/gorm"
)

type IRepository interface {
	Persist(loanBook model.LoanBook) (*model.LoanBook, error)
	Update(loanBook model.LoanBook) (*model.LoanBook, error)
	FindByUsers(loanBook model.LoanBook) (*model.LoanBook, error)
	FindByToUser(loanBook model.LoanBook) (*model.LoanBook, error)
	FindByBookAndStatus(bookId uint, status string) (*model.LoanBook, error)
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
	result := r.db.Create(&loanBook)

	if err := result.Error; err != nil {
		return nil, err
	}

	return &loanBook, nil
}

func (r Repository) Update(loanBook model.LoanBook) (*model.LoanBook, error) {
	result := r.db.Updates(&loanBook) //TODO update com status
	r.db.Model(&loanBook).Updates(&loanBook).Update("isActive", loanBook.Status)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &loanBook, nil
}

func (r Repository) FindByUsers(loanBook model.LoanBook) (*model.LoanBook, error) {
	result := r.db.Where("book_id = ? and from_user = ? and to_user = ? and status = ?",
						loanBook.BookID, loanBook.FromUser, loanBook.ToUser, loanBook.Status).
		First(&loanBook)

	if err := result.Error; err != nil {
		return nil, err
	}

	return &loanBook, nil
}

func (r Repository) FindByToUser(loanBook model.LoanBook) (*model.LoanBook, error) {
	result := r.db.Where("book_id = ? and to_user = ? and status = ?",
						loanBook.BookID, loanBook.ToUser, loanBook.Status).
		First(&loanBook)

	if err := result.Error; err != nil {
		return nil, err
	}

	return &loanBook, nil
}


func (r Repository) FindByBookAndStatus(bookId uint, status string) (*model.LoanBook, error) {
	var loanBook model.LoanBook
	result := r.db.Where("book_id = ? and status = ?",
		bookId, status).
		First(&loanBook)

	if err := result.Error; err != nil {
		return nil, err
	}

	return &loanBook, nil
}