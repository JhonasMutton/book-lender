package repository

import (
	"github.com/JhonasMutton/book-lender/pkg/repository/user"
	"github.com/google/wire"
)

var Set = wire.NewSet(user.NewRepository, wire.Bind(new(user.IRepository), new(*user.Repository)))
