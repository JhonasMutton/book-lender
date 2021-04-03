package book

import (
	"fmt"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type IRepository interface {
	Persist(book model.Book) (*model.Book, error)
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

func (r Repository) Persist(book model.Book) (*model.Book, error) {
	result := r.db.Create(&book)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &book, nil
}
