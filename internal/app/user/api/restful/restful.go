package restful

import (
	"github.com/blackhorseya/user-app/internal/pkg/infra/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// import swagger spec
	_ "github.com/blackhorseya/user-app/api/docs"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn() http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("api")
		{
			api.GET("readiness", func(c *gin.Context) {
				c.JSON(200, gin.H{"msg": "ok"})
			})
			api.GET("liveness", func(c *gin.Context) {
				c.JSON(200, gin.H{"msg": "ok"})
			})

			// open any environments can access swagger
			api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}

// ProviderSet is an apis provider set
var ProviderSet = wire.NewSet(
	CreateInitHandlerFn,
)
