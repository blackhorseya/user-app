package auth

import (
	"github.com/blackhorseya/user-app/internal/app/user/biz/auth"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	// import entity
	_ "github.com/blackhorseya/gocommon/pkg/er"
	_ "github.com/blackhorseya/gocommon/pkg/response"
)

type impl struct {
	logger *zap.Logger
	biz    auth.IBiz
}

// NewImpl return IHandler
func NewImpl(logger *zap.Logger, biz auth.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "auth.handler")),
		biz:    biz,
	}
}

// GetLoginURL
// @Summary Get login url
// @Description Get login url
// @Tags Auth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=string}
// @Failure 500 {object} er.APPError
// @Router /v1/auth/login [get]
func (i *impl) GetLoginURL(c *gin.Context) {
	// todo: 2022-03-01|05:40|Sean|impl me
	panic("implement me")
}

// Callback
// @Summary Callback after login
// @Description Callback after login
// @Tags Auth
// @Accept application/json
// @Produce application/json
// @Success 307 {string} string
// @Failure 500 {object} er.APPError
// @Router /v1/auth/callback [get]
func (i *impl) Callback(c *gin.Context) {
	// todo: 2022-03-01|05:40|Sean|impl me
	panic("implement me")
}

// Me
// @Summary Get me information
// @Description Get me information
// @Tags Auth
// @Accept application/json
// @Produce application/json
// @Security Bearer
// @Success 200 {object} response.Response{data=pb.Profile}
// @Failure 401 {object} er.APPError
// @Failure 403 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/auth/me [get]
func (i *impl) Me(c *gin.Context) {
	// todo: 2022-03-01|05:40|Sean|impl me
	panic("implement me")
}
