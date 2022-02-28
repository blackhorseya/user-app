package repo

import (
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
