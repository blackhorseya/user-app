//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/blackhorseya/user-app/internal/app/user/biz/auth/repo"
	"github.com/blackhorseya/user-app/internal/pkg/infra/authenticator"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateIBiz(logger *zap.Logger, repo repo.IRepo, auth authenticator.Authenticator) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
