// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package auth

import (
	"github.com/blackhorseya/user-app/internal/app/user/biz/auth/repo"
	"github.com/blackhorseya/user-app/internal/pkg/infra/authenticator"
	"github.com/blackhorseya/user-app/internal/pkg/infra/jwt"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func CreateIBiz(logger *zap.Logger, repo2 repo.IRepo, auth authenticator.Authenticator, jwt2 jwt.IJwt) (IBiz, error) {
	iBiz := NewImpl(logger, repo2, auth, jwt2)
	return iBiz, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)
