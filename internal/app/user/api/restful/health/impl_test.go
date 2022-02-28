package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/ginhttp"
	"github.com/blackhorseya/user-app/internal/app/user/biz/health/mocks"
	"github.com/blackhorseya/user-app/internal/pkg/entity/er"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type handlerSuite struct {
	suite.Suite
	r       *gin.Engine
	bizMock *mocks.IBiz
	handler IHandler
}

func (s *handlerSuite) SetupTest() {
	logger := zap.NewNop()

	gin.SetMode(gin.TestMode)
	s.r = gin.New()
	s.r.Use(ginhttp.AddContextx())
	s.r.Use(ginhttp.HandleError())

	s.bizMock = new(mocks.IBiz)
	handler, err := CreateIHandler(logger, s.bizMock)
	if err != nil {
		panic(err)
	}
	s.handler = handler
}

func (s *handlerSuite) TearDownTest() {
	s.bizMock.AssertExpectations(s.T())
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) Test_impl_Readiness() {
	s.r.GET("/api/readiness", s.handler.Readiness)

	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "readiness then error",
			args: args{mock: func() {
				s.bizMock.On("Readiness", mock.Anything).Return(er.ErrPingDatabase).Once()
			}},
			wantCode: 500,
		},
		{
			name: "readiness then success",
			args: args{mock: func() {
				s.bizMock.On("Readiness", mock.Anything).Return(nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := "/api/readiness"
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Readiness() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Liveness() {
	s.r.GET("/api/liveness", s.handler.Liveness)

	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "liveness then error",
			args: args{mock: func() {
				s.bizMock.On("Liveness", mock.Anything).Return(er.ErrPingDatabase).Once()
			}},
			wantCode: 500,
		},
		{
			name: "liveness then success",
			args: args{mock: func() {
				s.bizMock.On("Liveness", mock.Anything).Return(nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := "/api/liveness"
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Liveness() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}
