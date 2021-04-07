package user

import (
	"github.com/JhonasMutton/book-lender/pkg/log"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"gorm.io/gorm"
)

type IRepository interface {
	Persist(user model.User) (*model.User, error)
	Fetch() (*model.Users, error)
	FetchById(id uint) (*model.User, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Fetch() (*model.Users, error) {
	log.Logger.Debugf("Fetching users")
	var users model.Users
	result := r.db.Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *Repository) Persist(user model.User) (*model.User, error) {
	log.Logger.Debugf("Persisting user %x", user.Name)
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *Repository) FetchById(id uint) (*model.User, error) {
	log.Logger.Debugf("Fetching user by id %x", id)
	var user model.User
	result := r.db.Debug().Preload("Collection").
		Preload("LentBooks").
		Preload("BorrowedBooks").
		First(&user, id)  //Usar preloads pode ser oneroso, o ideal seria utilizar JOINS, porém despenderia mais tempo de trabalho
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
