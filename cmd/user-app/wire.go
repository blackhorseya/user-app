//go:build wireinject
// +build wireinject

package main

import (
	"github.com/blackhorseya/gocommon/pkg/config"
	"github.com/blackhorseya/gocommon/pkg/log"
	"github.com/blackhorseya/user-app/internal/app/user"
	"github.com/blackhorseya/user-app/internal/app/user/api/restful"
	"github.com/blackhorseya/user-app/internal/app/user/biz"
	"github.com/blackhorseya/user-app/internal/pkg/app"
	"github.com/blackhorseya/user-app/internal/pkg/infra/authenticator"
	"github.com/blackhorseya/user-app/internal/pkg/infra/databases"
	"github.com/blackhorseya/user-app/internal/pkg/infra/jwt"
	"github.com/blackhorseya/user-app/internal/pkg/infra/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	user.ProviderSet,
	config.ProviderSet,
	log.ProviderSet,
	http.ProviderSet,
	databases.ProviderSet,
	jwt.ProviderSet,
	authenticator.ProviderSet,
	restful.ProviderSet,
	biz.ProviderSet,
)

func CreateApp(path string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
