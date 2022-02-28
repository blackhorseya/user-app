package repo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/google/wire"
)

// IRepo declare repository functions
//go:generate mockery --name=IRepo
type IRepo interface {
	PingDatabase(ctx contextx.Contextx) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
