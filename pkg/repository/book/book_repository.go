package book

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"gorm.io/gorm"
)

type IRepository interface {
	Persist(book model.Book) (*model.Book, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Persist(book model.Book) (*model.Book, error) {
	result := r.db.Create(&book)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &book, nil
}
