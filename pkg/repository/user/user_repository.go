package user

import (
	"fmt"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type IRepository interface {
	FetchUsers() (*model.Users, error)
	PersistUser(user model.User) (*model.User, error)
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

	err = db.AutoMigrate(&model.User{}) //TODO Tirar migrate daqui
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&model.LoanBooks{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&model.Book{})
	if err != nil {
		panic(err.Error())
	}

	return &Repository{
		db: db,
	}
}

func (r *Repository) FetchUsers() (*model.Users, error) {
	var users model.Users
	result := r.db.Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *Repository) PersistUser(user model.User) (*model.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
