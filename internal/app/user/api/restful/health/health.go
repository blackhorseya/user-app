package health

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare health api handlers
type IHandler interface {
	Readiness(c *gin.Context)

	Liveness(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
