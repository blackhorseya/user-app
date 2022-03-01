// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package repo

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func CreateIRepo(logger *zap.Logger, client *mongo.Client) (IRepo, error) {
	iRepo := NewImpl(logger, client)
	return iRepo, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)