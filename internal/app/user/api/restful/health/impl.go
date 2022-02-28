package health

import (
	"net/http"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/user-app/internal/app/user/biz/health"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	// import entities
	_ "github.com/blackhorseya/gocommon/pkg/er"
)

type impl struct {
	logger *zap.Logger
	biz    health.IBiz
}

// NewImpl return IHandler
func NewImpl(logger *zap.Logger, biz health.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "health.restful")),
		biz:    biz,
	}
}

// Readiness to know when an application is ready to start accepting traffic
// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=string}
// @Failure 500 {object} er.APPError
// @Router /readiness [get]
func (i *impl) Readiness(c *gin.Context) {
	ctx := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)

	err := i.biz.Readiness(ctx)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData("success"))
}

// Liveness to know when to restart an application
// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=string}
// @Failure 500 {object} er.APPError
// @Router /liveness [get]
func (i *impl) Liveness(c *gin.Context) {
	ctx := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)

	err := i.biz.Liveness(ctx)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData("success"))
}
