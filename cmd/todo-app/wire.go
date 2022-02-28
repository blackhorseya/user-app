//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

var providerSet = wire.NewSet()

func CreateApp(path string) {
	panic(wire.Build(providerSet))
}
