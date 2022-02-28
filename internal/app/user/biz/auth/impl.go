package auth

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/user-app/internal/app/user/biz/auth/repo"
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	"github.com/blackhorseya/user-app/internal/pkg/infra/authenticator"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	auth   authenticator.Authenticator
	repo   repo.IRepo
}

// NewImpl return IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo, auth authenticator.Authenticator) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "auth.biz")),
		repo:   repo,
		auth:   auth,
	}
}

func (i *impl) GetLoginURL(ctx contextx.Contextx, state string) string {
	return i.auth.AuthCodeURL(ctx, state)
}

func (i *impl) Callback(ctx contextx.Contextx, code string) (info *user.Profile, err error) {
	// todo: 2022-03-01|05:26|Sean|impl me
	panic("implement me")
}

func (i *impl) GetUserByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error) {
	// todo: 2022-03-01|05:26|Sean|impl me
	panic("implement me")
}

func (i *impl) HasPermission(ctx contextx.Contextx, token, obj, act string) (bool, error) {
	// todo: 2022-03-01|05:26|Sean|impl me
	panic("implement me")
}
