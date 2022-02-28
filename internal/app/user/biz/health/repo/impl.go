package repo

import (
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	client *mongo.Client
}

// NewImpl return IRepo
func NewImpl(logger *zap.Logger, client *mongo.Client) IRepo {
	return &impl{
		logger: logger.With(zap.String("type", "health.repo")),
		client: client,
	}
}

func (i *impl) PingDatabase(ctx contextx.Contextx) error {
	timeout, cancel := contextx.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := i.client.Ping(timeout, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}
