package databases

import (
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// Options is configuration of database
type Options struct {
	URL string `json:"url" yaml:"url"`
}

// NewOptions serve caller to create an Options
func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	logger.Info("load database options success", zap.String("url", o.URL))

	return o, err
}

// NewMongo init a mongo client
func NewMongo(o *Options) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(o.URL))
	if err != nil {
		return nil, err
	}

	ctx, cancel := contextx.WithTimeout(contextx.Background(), 5*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewOptions, NewMongo)
