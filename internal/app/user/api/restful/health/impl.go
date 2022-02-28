package health

import (
	"net/http"

	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	// import entities
	_ "github.com/blackhorseya/gocommon/pkg/er"
)

type impl struct {
	logger *zap.Logger
}

// NewImpl return IHandler
func NewImpl(logger *zap.Logger) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "health.restful")),
	}
}

// Readiness to know when an application is ready to start accepting traffic
// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} er.APPError
// @Router /readiness [get]
func (i *impl) Readiness(c *gin.Context) {
	c.JSON(http.StatusOK, response.OK)
}

// Liveness to know when to restart an application
// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} er.APPError
// @Router /liveness [get]
func (i *impl) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, response.OK)
}
