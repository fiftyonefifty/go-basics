//+build wireinject

package main

import (
	"github.com/google/wire"
)

func CreateConcatService() *ConcatService {
	panic(wire.Build(
		Logger{},
		NewHttpClient,
		NewConcatService,
	))
}
