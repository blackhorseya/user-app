//go:build wireinject
// +build wireinject

package health

import (
	"github.com/blackhorseya/user-app/internal/app/user/biz/health/repo"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateIBiz(logger *zap.Logger, repo repo.IRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
