package handler

import (
	"github.com/JhonasMutton/book-lender/pkg/api/handler/book"
	"github.com/JhonasMutton/book-lender/pkg/api/handler/lend"
	"github.com/JhonasMutton/book-lender/pkg/api/handler/user"
	"github.com/google/wire"
)

var Set =  wire.NewSet(user.NewHandler, book.NewHandler, lend.NewHandler)
