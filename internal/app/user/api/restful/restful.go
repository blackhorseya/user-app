package restful

import (
	"github.com/blackhorseya/user-app/internal/app/user/api/restful/health"
	"github.com/blackhorseya/user-app/internal/pkg/infra/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// import swagger spec
	_ "github.com/blackhorseya/user-app/api/docs"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(healthH health.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("api")
		{
			api.GET("readiness", healthH.Readiness)
			api.GET("liveness", healthH.Liveness)

			// open any environments can access swagger
			api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}

// ProviderSet is an apis provider set
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	CreateInitHandlerFn,
)
