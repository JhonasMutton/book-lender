package lend

import (
	"fmt"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type IRepository interface {
	Persist(loanBook model.LoanBook) (*model.LoanBook, error)
	Update(loanBook model.LoanBook) (*model.LoanBook, error)
	Find(loanBook model.LoanBook) (*model.LoanBook, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=60s&readTimeout=60s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

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

	if err := result.Error; err != nil {
		return nil, err
	}

	return &loanBook, nil
}

func (r Repository) Find(loanBook model.LoanBook) (*model.LoanBook, error) {
	result := r.db.First(&loanBook) //TODO FIND COM STatus incluso

	if err := result.Error; err != nil {
		return nil, err
	}

	return &loanBook, nil
}
