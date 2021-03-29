package config

import "github.com/google/wire"

var ConfigSet = wire.NewSet(NewHandler)
