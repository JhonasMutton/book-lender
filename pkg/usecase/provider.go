package usecase

import (
	"github.com/JhonasMutton/book-lender/pkg/usecase/book"
	"github.com/JhonasMutton/book-lender/pkg/usecase/lend"
	"github.com/JhonasMutton/book-lender/pkg/usecase/user"
	"github.com/google/wire"
)

var Set = wire.NewSet(user.NewUseCase, wire.Bind(new(user.IUseCase), new(*user.UseCase)),
						book.NewUseCase, wire.Bind(new(book.IUseCase), new(*book.UseCase)),
							lend.NewUseCase, wire.Bind(new(lend.IUseCase), new(*lend.UseCase)))
