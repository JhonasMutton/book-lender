package pkg

import (
	"github.com/JhonasMutton/book-lender/pkg/api/config"
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(config.ConfigSet, ApplicationSet)
