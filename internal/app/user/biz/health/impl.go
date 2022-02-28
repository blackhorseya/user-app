package health

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/user-app/internal/app/user/biz/health/repo"
	"github.com/blackhorseya/user-app/internal/pkg/entity/er"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
}

// NewImpl return IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "health.biz")),
		repo:   repo,
	}
}

func (i *impl) Readiness(ctx contextx.Contextx) error {
	err := i.repo.PingDatabase(ctx)
	if err != nil {
		i.logger.Error(er.ErrPingDatabase.Error(), zap.Error(err))
		return er.ErrPingDatabase
	}

	return nil
}

func (i *impl) Liveness(ctx contextx.Contextx) error {
	err := i.repo.PingDatabase(ctx)
	if err != nil {
		i.logger.Error(er.ErrPingDatabase.Error(), zap.Error(err))
		return er.ErrPingDatabase
	}

	return nil
}
