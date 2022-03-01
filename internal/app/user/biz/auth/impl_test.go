package auth

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	repoMocks "github.com/blackhorseya/user-app/internal/app/user/biz/auth/repo/mocks"
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	authMocks "github.com/blackhorseya/user-app/internal/pkg/infra/authenticator/mocks"
	jwtMocks "github.com/blackhorseya/user-app/internal/pkg/infra/jwt/mocks"
	"github.com/blackhorseya/user-app/test/testdata"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type suiteBiz struct {
	suite.Suite
	mockRepo *repoMocks.IRepo
	mockAuth *authMocks.Authenticator
	mockJwt  *jwtMocks.IJwt
	biz      IBiz
}

func (s *suiteBiz) SetupTest() {
	logger := zap.NewNop()

	s.mockRepo = new(repoMocks.IRepo)
	s.mockAuth = new(authMocks.Authenticator)
	s.mockJwt = new(jwtMocks.IJwt)
	biz, err := CreateIBiz(logger, s.mockRepo, s.mockAuth, s.mockJwt)
	if err != nil {
		panic(err)
	}
	s.biz = biz
}

func (s *suiteBiz) TearDownTest() {
	s.mockRepo.AssertExpectations(s.T())
	s.mockAuth.AssertExpectations(s.T())
}

func TestSuiteBiz(t *testing.T) {
	suite.Run(t, new(suiteBiz))
}

func (s *suiteBiz) Test_impl_GetLoginURL() {
	type args struct {
		state string
		mock  func()
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get login url then success",
			args: args{state: "123", mock: func() {
				s.mockAuth.On("AuthCodeURL", mock.Anything, "123").Return("url").Once()
			}},
			want: "url",
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if got := s.biz.GetLoginURL(contextx.Background(), tt.args.state); got != tt.want {
				t.Errorf("GetLoginURL() = %v, want %v", got, tt.want)
			}

			s.TearDownTest()
		})
	}
}

func (s *suiteBiz) Test_impl_Callback() {
	type args struct {
		code string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name: "exchange code to token then error",
			args: args{code: "code", mock: func() {
				s.mockAuth.On("Exchange", mock.Anything, "code").Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "exchange token to verify id token then error",
			args: args{code: "code", mock: func() {
				s.mockAuth.On("Exchange", mock.Anything, "code").Return(&oauth2.Token{}, nil).Once()
				s.mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get claim from id token then error",
			args: args{code: "code", mock: func() {
				s.mockAuth.On("Exchange", mock.Anything, "code").Return(&oauth2.Token{}, nil).Once()
				s.mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&oidc.IDToken{}, nil).Once()
				s.mockAuth.On("Claims", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get user by open id then error",
			args: args{code: "code", mock: func() {
				s.mockAuth.On("Exchange", mock.Anything, "code").Return(&oauth2.Token{}, nil).Once()
				s.mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&oidc.IDToken{}, nil).Once()
				s.mockAuth.On("Claims", mock.Anything).Return(testdata.Claims1, nil).Once()

				s.mockRepo.On("GetUserByOpenID", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "if exists then update user then error",
			args: args{code: "code", mock: func() {
				s.mockAuth.On("Exchange", mock.Anything, "code").Return(&oauth2.Token{}, nil).Once()
				s.mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&oidc.IDToken{}, nil).Once()
				s.mockAuth.On("Claims", mock.Anything).Return(testdata.Claims1, nil).Once()

				s.mockRepo.On("GetUserByOpenID", mock.Anything, mock.Anything, mock.Anything).Return(testdata.User1, nil).Once()

				s.mockJwt.On("NewToken", mock.Anything).Return("token", nil).Once()
				s.mockRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "update user then success",
			args: args{code: "code", mock: func() {
				s.mockAuth.On("Exchange", mock.Anything, "code").Return(&oauth2.Token{}, nil).Once()
				s.mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&oidc.IDToken{}, nil).Once()
				s.mockAuth.On("Claims", mock.Anything).Return(testdata.Claims1, nil).Once()

				s.mockRepo.On("GetUserByOpenID", mock.Anything, mock.Anything, mock.Anything).Return(testdata.User1, nil).Once()

				s.mockJwt.On("NewToken", mock.Anything).Return("token", nil).Once()
				s.mockRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(testdata.User1, nil).Once()
			}},
			wantInfo: testdata.User1,
			wantErr:  false,
		},
		{
			name: "if not exists then create new user then error",
			args: args{code: "code", mock: func() {
				s.mockAuth.On("Exchange", mock.Anything, "code").Return(&oauth2.Token{}, nil).Once()
				s.mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&oidc.IDToken{}, nil).Once()
				s.mockAuth.On("Claims", mock.Anything).Return(testdata.Claims1, nil).Once()

				s.mockRepo.On("GetUserByOpenID", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()

				s.mockJwt.On("NewToken", mock.Anything).Return("token", nil).Once()
				s.mockRepo.On("RegisterUser", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "register user then success",
			args: args{code: "code", mock: func() {
				s.mockAuth.On("Exchange", mock.Anything, "code").Return(&oauth2.Token{}, nil).Once()
				s.mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&oidc.IDToken{}, nil).Once()
				s.mockAuth.On("Claims", mock.Anything).Return(testdata.Claims1, nil).Once()

				s.mockRepo.On("GetUserByOpenID", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()

				s.mockJwt.On("NewToken", mock.Anything).Return("token", nil).Once()
				s.mockRepo.On("RegisterUser", mock.Anything, mock.Anything).Return(testdata.User1, nil).Once()
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

			gotInfo, err := s.biz.Callback(contextx.Background(), tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Callback() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Callback() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

func (s *suiteBiz) Test_impl_GetUserByToken() {
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
			name: "get by token then error",
			args: args{token: testdata.User1.Token, mock: func() {
				s.mockRepo.On("GetUserByToken", mock.Anything, testdata.User1.Token).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by token then success",
			args: args{token: testdata.User1.Token, mock: func() {
				s.mockRepo.On("GetUserByToken", mock.Anything, testdata.User1.Token).Return(testdata.User1, nil).Once()
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

			gotInfo, err := s.biz.GetUserByToken(contextx.Background(), tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetUserByToken() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}
