package user

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/user"
	"time"
)

type IUseCase interface {
	CreateUser(basicUser model.BasicUser) (*model.User, error)
	FindUsers() (*model.Users, error)
}

type UseCase struct {
	userRepository user.IRepository
}

func NewUseCase(userRepository user.IRepository) *UseCase {
	return &UseCase{userRepository: userRepository}
}

func (u UseCase) CreateUser(basicUser model.BasicUser) (*model.User, error) {
	us := model.User{
		BasicUser: basicUser,
		CreatedAt: time.Now(),
	}

	persisted, err := u.userRepository.PersistUser(us)
	if err != nil {
		return nil, err //TODO Handle errors
	}

	return persisted, nil
}

func (u UseCase) FindUsers() (*model.Users, error) {
	users, err := u.userRepository.FetchUsers()
	if err != nil {
		return nil, err //TODO Handle errors
	}

	return users, nil
}
