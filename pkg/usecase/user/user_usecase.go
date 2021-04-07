package user

import (
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/JhonasMutton/book-lender/pkg/repository/user"
	"github.com/JhonasMutton/book-lender/pkg/validate"
	"strconv"
)

type IUseCase interface {
	Create(userDto model.UserDto) (*model.User, error)
	Find() (*model.Users, error)
	FindById(id string) (*model.User, error)
}

type UseCase struct {
	userRepository user.IRepository
	validator      *validate.Validator
}

func NewUseCase(userRepository user.IRepository, validate *validate.Validator) *UseCase {
	return &UseCase{userRepository: userRepository, validator: validate}
}

func (u UseCase) Create(userDto model.UserDto) (*model.User, error) {
	if err := u.validator.Validate(userDto); err != nil {
		return nil, errors.WrapWithMessage(errors.ErrInvalidPayload, err.Error())
	}

	userModel := userDto.ToModel()

	persisted, err := u.userRepository.Persist(userModel)
	if err != nil {
		return nil, errors.BuildError(err)
	}

	return persisted, nil
}

func (u UseCase) Find() (*model.Users, error) {
	users, err := u.userRepository.Fetch()
	if err != nil {
		return nil, errors.BuildError(err)
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
		return nil, errors.BuildError(err)
	}

	return users, nil
}
