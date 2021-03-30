//+build wireinject

package main

import (
	"github.com/JhonasMutton/book-lender/pkg"
	"github.com/google/wire"
)

func SetupApplication() pkg.Application {
	wire.Build(pkg.DependencySet)
	return pkg.Application{}
}
