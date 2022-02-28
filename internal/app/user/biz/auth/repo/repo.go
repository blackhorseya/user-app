package repo

import "github.com/google/wire"

// IRepo declare auth repository functions
//go:generate mockery --name=IRepo
type IRepo interface {
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
