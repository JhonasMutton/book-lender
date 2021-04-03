package user

import (
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/user"
	"strconv"
)

type IUseCase interface {
	Create(userDto model.UserDto) (*model.User, error)
	Find() (*model.Users, error)
	FindById(id string) (*model.User, error)
}

type UseCase struct {
	userRepository user.IRepository
}

func NewUseCase(userRepository user.IRepository) *UseCase {
	return &UseCase{userRepository: userRepository}
}

func (u UseCase) Create(userDto model.UserDto) (*model.User, error) {
	//TODO VALIDATE FIELDS
	userModel := userDto.ToModel()

	persisted, err := u.userRepository.Persist(userModel)
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
