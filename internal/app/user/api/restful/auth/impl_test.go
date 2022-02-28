package auth

import (
	"encoding/gob"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/ginhttp"
	bizMocks "github.com/blackhorseya/user-app/internal/app/user/biz/auth/mocks"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteHandler struct {
	suite.Suite
	r       *gin.Engine
	mockBiz *bizMocks.IBiz
	handler IHandler
}

func (s *suiteHandler) SetupTest() {
	logger := zap.NewNop()

	gin.SetMode(gin.TestMode)
	s.r = gin.New()
	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	s.r.Use(sessions.Sessions("auth-session", store))
	s.r.Use(ginhttp.AddContextx())
	s.r.Use(ginhttp.HandleError())

	s.mockBiz = new(bizMocks.IBiz)
	handler, err := CreateIHandler(logger, s.mockBiz)
	if err != nil {
		panic(err)
	}
	s.handler = handler
}

func (s *suiteHandler) TearDownTest() {
	s.mockBiz.AssertExpectations(s.T())
}

func TestSuiteHandler(t *testing.T) {
	suite.Run(t, new(suiteHandler))
}

func (s *suiteHandler) Test_impl_GetLoginURL() {
	s.r.GET("/api/v1/auth/login", s.handler.GetLoginURL)

	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "get login url then 200",
			args: args{mock: func() {
				s.mockBiz.On("GetLoginURL", mock.Anything, mock.Anything).Return("url").Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := "/api/v1/auth/login"
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "GetLoginURL() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}
