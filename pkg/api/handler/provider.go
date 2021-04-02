package handler

import (
	"github.com/JhonasMutton/book-lender/pkg/api/handler/user"
	"github.com/google/wire"
)

var Set =  wire.NewSet(user.NewHandler)
