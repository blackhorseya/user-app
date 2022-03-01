package repo

import (
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	dbName = "side"

	collName = "users"
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
	timeout, cancel := contextx.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var ret user.Profile
	filter := bson.M{"open_ids." + provider: id}
	coll := i.client.Database(dbName).Collection(collName)
	err = coll.FindOne(timeout, filter).Decode(&ret)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &ret, nil
}

func (i *impl) GetUserByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var ret user.Profile
	filter := bson.M{"token": token}
	coll := i.client.Database(dbName).Collection(collName)
	err = coll.FindOne(timeout, filter).Decode(&ret)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &ret, nil
}

func (i *impl) RegisterUser(ctx contextx.Contextx, newUser *user.Profile) (info *user.Profile, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	now := time.Now()
	newUser.CreatedAt = primitive.NewDateTimeFromTime(now)
	newUser.UpdatedAt = primitive.NewDateTimeFromTime(now)

	coll := i.client.Database(dbName).Collection(collName)
	res, err := coll.InsertOne(timeout, newUser)
	if err != nil {
		return nil, err
	}

	newUser.ID = res.InsertedID.(primitive.ObjectID)

	return newUser, nil
}

func (i *impl) UpdateUser(ctx contextx.Contextx, newUser *user.Profile) (info *user.Profile, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	newUser.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	var ret user.Profile
	filter := bson.M{"_id": newUser.ID}
	opts := options.FindOneAndReplace().SetUpsert(false).SetReturnDocument(options.After)
	coll := i.client.Database(dbName).Collection(collName)
	err = coll.FindOneAndReplace(timeout, filter, newUser, opts).Decode(&ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
