//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/blackhorseya/user-app/internal/app/user/biz/auth"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateIHandler(logger *zap.Logger, biz auth.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
