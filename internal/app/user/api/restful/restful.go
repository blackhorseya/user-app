package restful

import (
	"github.com/blackhorseya/user-app/internal/app/user/api/restful/auth"
	"github.com/blackhorseya/user-app/internal/app/user/api/restful/health"
	authB "github.com/blackhorseya/user-app/internal/app/user/biz/auth"
	"github.com/blackhorseya/user-app/internal/pkg/infra/transports/http"
	"github.com/blackhorseya/user-app/internal/pkg/infra/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// import swagger spec
	_ "github.com/blackhorseya/user-app/api/docs"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(
	authBiz authB.IBiz,
	healthH health.IHandler,
	authH auth.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("api")
		{
			api.GET("readiness", healthH.Readiness)
			api.GET("liveness", healthH.Liveness)

			// open any environments can access swagger
			api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

			v1 := api.Group("v1")
			{
				authG := v1.Group("auth")
				{
					authG.GET("login", authH.GetLoginURL)
					authG.GET("callback", authH.Callback)
					authG.GET("me", middlewares.RequiredAuth(authBiz), authH.Me)
				}
			}
		}
	}
}

// ProviderSet is an apis provider set
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	auth.ProviderSet,
	CreateInitHandlerFn,
)
