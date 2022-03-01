package middlewares

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/gocommon/pkg/utils/tokenutil"
	"github.com/blackhorseya/user-app/internal/app/user/biz/auth"
	"github.com/gin-gonic/gin"
)

// RequiredAuth required authentication in header
func RequiredAuth(auth auth.IBiz) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := tokenutil.GetBearerToken(c)
		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}

		ctx := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
		user, err := auth.GetUserByToken(ctx, token)
		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}

		c.Set("ctx", contextx.WithValue(ctx, "user", user))

		c.Next()
	}
}
