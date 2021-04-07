package database

import "github.com/google/wire"

var Set = wire.NewSet(NewDatabaseConnection)
