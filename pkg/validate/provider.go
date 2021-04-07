package validate

import "github.com/google/wire"

var Set =  wire.NewSet(NewValidator)
