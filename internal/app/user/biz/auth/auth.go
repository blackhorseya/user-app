package auth

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/user-app/internal/app/user/biz/auth/repo"
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	"github.com/google/wire"
)

// IBiz declare auth business functions
//go:generate mockery --name=IBiz
type IBiz interface {
	GetLoginURL(ctx contextx.Contextx, state string) string

	Callback(ctx contextx.Contextx, code string) (info *user.Profile, err error)

	GetUserByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error)

	HasPermission(ctx contextx.Contextx, token, obj, act string) (bool, error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
