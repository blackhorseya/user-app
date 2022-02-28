//go:build wireinject
// +build wireinject

package health

import (
	"github.com/blackhorseya/user-app/internal/app/user/biz/health"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateIHandler(logger *zap.Logger, biz health.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
