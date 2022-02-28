package repo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	client *mongo.Client
}

// NewImpl return IRepo
func NewImpl(logger *zap.Logger, client *mongo.Client) IRepo {
	return &impl{
		logger: logger.With(zap.String("type", "auth.repo")),
		client: client,
	}
}

func (i *impl) GetUserByOpenID(ctx contextx.Contextx, provider, id string) (info *user.Profile, err error) {
	// todo: 2022-03-01|06:26|Sean|impl me
	panic("implement me")
}

func (i *impl) GetUserByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error) {
	// todo: 2022-03-01|06:26|Sean|impl me
	panic("implement me")
}

func (i *impl) RegisterUser(ctx contextx.Contextx, newUser *user.Profile) (info *user.Profile, err error) {
	// todo: 2022-03-01|06:26|Sean|impl me
	panic("implement me")
}

func (i *impl) UpdateUser(ctx contextx.Contextx, newUser *user.Profile) (info *user.Profile, err error) {
	// todo: 2022-03-01|06:26|Sean|impl me
	panic("implement me")
}
