package pkg

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
)

var ApplicationSet = wire.NewSet(NewApplication, validator.New)
