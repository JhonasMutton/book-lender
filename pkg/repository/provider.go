package repository

import (
	"github.com/JhonasMutton/book-lender/pkg/repository/book"
	"github.com/JhonasMutton/book-lender/pkg/repository/lend"
	"github.com/JhonasMutton/book-lender/pkg/repository/user"
	"github.com/google/wire"
)

var Set = wire.NewSet(user.NewRepository, wire.Bind(new(user.IRepository), new(*user.Repository)),
						book.NewRepository, wire.Bind(new(book.IRepository), new(*book.Repository)),
							lend.NewRepository, wire.Bind(new(lend.IRepository), new(*lend.IRepository)))
