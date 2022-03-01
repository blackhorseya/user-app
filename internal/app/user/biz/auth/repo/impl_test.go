package repo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	"github.com/blackhorseya/user-app/test/testdata"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type suiteRepo struct {
	suite.Suite
	pool     *dockertest.Pool
	resource *dockertest.Resource
	client   *mongo.Client
	repo     IRepo
}

func (s *suiteRepo) SetupTest() {
	logger := zap.NewNop()

	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}
	s.pool = pool

	resource, err := pool.Run("mongo", "4.4.10", nil)
	if err != nil {
		panic(err)
	}
	s.resource = resource

	err = pool.Retry(func() error {
		uri := fmt.Sprintf("mongodb://localhost:%s/", resource.GetPort("27017/tcp"))
		s.client, err = mongo.Connect(contextx.Background(), options.Client().ApplyURI(uri))
		if err != nil {
			return err
		}

		return s.client.Ping(contextx.Background(), readpref.Primary())
	})
	if err != nil {
		panic(err)
	}

	repo, err := CreateIRepo(logger, s.client)
	if err != nil {
		panic(err)
	}
	s.repo = repo
}

func (s *suiteRepo) TearDownTest() {
	_ = s.client.Disconnect(contextx.Background())
	_ = s.pool.Purge(s.resource)
}

func TestSuiteRepo(t *testing.T) {
	suite.Run(t, new(suiteRepo))
}

func (s *suiteRepo) Test_impl_GetUserByOpenID() {
	type args struct {
		provider string
		id       string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "not found then nil",
			args:     args{provider: "line", id: "line"},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "get by open id then success",
			args: args{provider: "line", id: "line", mock: func() {
				_, _ = s.client.Database(dbName).Collection(collName).InsertOne(contextx.Background(), testdata.User1)
			}},
			wantInfo: testdata.User1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.GetUserByOpenID(contextx.Background(), tt.args.provider, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByOpenID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetUserByOpenID() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			_, _ = s.client.Database(dbName).Collection(collName).DeleteMany(contextx.Background(), bson.M{})
		})
	}
}

func (s *suiteRepo) Test_impl_GetUserByToken() {
	type args struct {
		token string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "get by token then not found",
			args:     args{token: "token"},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "get by token then not found",
			args: args{token: "token", mock: func() {
				_, _ = s.client.Database(dbName).Collection(collName).InsertOne(contextx.Background(), testdata.User1)
			}},
			wantInfo: testdata.User1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.GetUserByToken(contextx.Background(), tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetUserByToken() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			_, _ = s.client.Database(dbName).Collection(collName).DeleteMany(contextx.Background(), bson.M{})
		})
	}
}

func (s *suiteRepo) Test_impl_RegisterUser() {
	type args struct {
		newUser *user.Profile
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name: "duplication then error",
			args: args{newUser: testdata.User1, mock: func() {
				_, _ = s.client.Database(dbName).Collection(collName).InsertOne(contextx.Background(), testdata.User1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "register user then success",
			args:     args{newUser: testdata.User1},
			wantInfo: testdata.User1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.RegisterUser(contextx.Background(), tt.args.newUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("RegisterUser() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			_, _ = s.client.Database(dbName).Collection(collName).DeleteMany(contextx.Background(), bson.M{})
		})
	}
}

func (s *suiteRepo) Test_impl_UpdateUser() {
	type args struct {
		newUser *user.Profile
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "update then not found",
			args:     args{newUser: testdata.User1},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "update then not found",
			args: args{newUser: testdata.User1, mock: func() {
				_, _ = s.client.Database(dbName).Collection(collName).InsertOne(contextx.Background(), testdata.User1)
			}},
			wantInfo: testdata.User1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.UpdateUser(contextx.Background(), tt.args.newUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("UpdateUser() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			_, _ = s.client.Database(dbName).Collection(collName).DeleteMany(contextx.Background(), bson.M{})
		})
	}
}
