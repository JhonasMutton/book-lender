package usecase

import (
	"github.com/JhonasMutton/book-lender/pkg/usecase/user"
	"github.com/google/wire"
)

var Set = wire.NewSet(user.NewUseCase, wire.Bind(new(user.IUseCase), new(*user.UseCase)))
