package pkg

import (
	"github.com/JhonasMutton/book-lender/pkg/api/config"
	"github.com/JhonasMutton/book-lender/pkg/api/handler"
	"github.com/JhonasMutton/book-lender/pkg/repository"
	"github.com/JhonasMutton/book-lender/pkg/usecase"
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(config.Set, repository.Set, usecase.Set, handler.Set, ApplicationSet)
