package repo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	"github.com/google/wire"
)

// IRepo declare auth repository functions
//go:generate mockery --name=IRepo
type IRepo interface {
	GetUserByOpenID(ctx contextx.Contextx, provider, id string) (info *user.Profile, err error)

	GetUserByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error)

	RegisterUser(ctx contextx.Contextx, newUser *user.Profile) (info *user.Profile, err error)

	UpdateUser(ctx contextx.Contextx, newUser *user.Profile) (info *user.Profile, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
