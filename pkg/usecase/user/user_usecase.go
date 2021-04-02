package user

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/user"
	"strconv"
	"time"
)

type IUseCase interface {
	Create(basicUser model.BasicUser) (*model.User, error)
	Find() (*model.Users, error)
	FindById(id string) (*model.User, error)
}

type UseCase struct {
	userRepository user.IRepository
}

func NewUseCase(userRepository user.IRepository) *UseCase {
	return &UseCase{userRepository: userRepository}
}

func (u UseCase) Create(basicUser model.BasicUser) (*model.User, error) {
	us := model.User{
		BasicUser: basicUser,
		CreatedAt: time.Now(),
	}

	persisted, err := u.userRepository.Persist(us)
	if err != nil {
		return nil, err //TODO Handle errors
	}

	return persisted, nil
}

func (u UseCase) Find() (*model.Users, error) {
	users, err := u.userRepository.Fetch()
	if err != nil {
		return nil, err //TODO Handle errors
	}

	return users, nil
}

func (u UseCase) FindById(id string) (*model.User, error) {
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}

	idUint := uint(id64)

	users, err := u.userRepository.FetchById(idUint)
	if err != nil {
		return nil, err //TODO Handle errors
	}

	return users, nil
}
