package health

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/user-app/internal/app/user/biz/health/repo"
	"github.com/google/wire"
)

// IBiz declare health business functions
//go:generate mockery --name=IBiz
type IBiz interface {
	Readiness(ctx contextx.Contextx) error

	Liveness(ctx contextx.Contextx) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
