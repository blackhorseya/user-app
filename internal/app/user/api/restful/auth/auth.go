package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare auth handlers
type IHandler interface {
	GetLoginURL(c *gin.Context)

	Callback(c *gin.Context)

	Me(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
