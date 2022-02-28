package health

import (
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/user-app/internal/app/user/biz/health/repo/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type bizSuite struct {
	suite.Suite
	repoMock *mocks.IRepo
	biz      IBiz
}

func (s *bizSuite) SetupTest() {
	logger := zap.NewNop()

	s.repoMock = new(mocks.IRepo)
	biz, err := CreateIBiz(logger, s.repoMock)
	if err != nil {
		panic(err)
	}
	s.biz = biz
}

func (s *bizSuite) TearDownTest() {
	s.repoMock.AssertExpectations(s.T())
}

func TestBizSuite(t *testing.T) {
	suite.Run(t, new(bizSuite))
}

func (s *bizSuite) Test_impl_Readiness() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ping then error",
			args: args{mock: func() {
				s.repoMock.On("PingDatabase", mock.Anything).Return(errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "ping then success",
			args: args{mock: func() {
				s.repoMock.On("PingDatabase", mock.Anything).Return(nil).Once()
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.Readiness(contextx.Background()); (err != nil) != tt.wantErr {
				t.Errorf("Readiness() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_Liveness() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ping then error",
			args: args{mock: func() {
				s.repoMock.On("PingDatabase", mock.Anything).Return(errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "ping then success",
			args: args{mock: func() {
				s.repoMock.On("PingDatabase", mock.Anything).Return(nil).Once()
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.Liveness(contextx.Background()); (err != nil) != tt.wantErr {
				t.Errorf("Liveness() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}
