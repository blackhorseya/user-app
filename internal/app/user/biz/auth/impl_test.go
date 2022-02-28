package auth

import (
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	repoMocks "github.com/blackhorseya/user-app/internal/app/user/biz/auth/repo/mocks"
	authMocks "github.com/blackhorseya/user-app/internal/pkg/infra/authenticator/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteBiz struct {
	suite.Suite
	mockRepo *repoMocks.IRepo
	mockAuth *authMocks.Authenticator
	biz      IBiz
}

func (s *suiteBiz) SetupTest() {
	logger := zap.NewNop()

	s.mockRepo = new(repoMocks.IRepo)
	s.mockAuth = new(authMocks.Authenticator)
	biz, err := CreateIBiz(logger, s.mockRepo, s.mockAuth)
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
