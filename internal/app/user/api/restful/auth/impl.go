package auth

import (
	"encoding/base64"
	"net/http"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/gocommon/pkg/utils/randutil"
	"github.com/blackhorseya/user-app/internal/app/user/biz/auth"
	cer "github.com/blackhorseya/user-app/internal/pkg/entity/er"
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	// import entity
	_ "github.com/blackhorseya/gocommon/pkg/er"
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
// @Param redirect query string true "redirect url after login success"
// @Success 200 {object} response.Response{data=string}
// @Failure 500 {object} er.APPError
// @Router /v1/auth/login [get]
func (i *impl) GetLoginURL(c *gin.Context) {
	ctx := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)

	state := base64.StdEncoding.EncodeToString([]byte(randutil.String(8)))

	session := sessions.Default(c)
	session.Set("state", state)
	session.Set("redirect", c.Query("redirect"))
	err := session.Save()
	if err != nil {
		_ = c.Error(err)
		return
	}

	ret := i.biz.GetLoginURL(ctx, state)

	c.JSON(http.StatusOK, response.OK.WithData(ret))
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
	ctx := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)

	session := sessions.Default(c)
	if c.Query("state") != session.Get("state") {
		_ = c.Error(cer.ErrInvalidState)
		return
	}

	ret, err := i.biz.Callback(ctx, c.Query("code"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	redirect := session.Get("redirect").(string)

	c.Redirect(http.StatusTemporaryRedirect, redirect+"?token="+ret.Token)
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
	ctx := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)

	ret := ctx.Value("user").(*user.Profile)

	c.JSON(http.StatusOK, response.OK.WithData(user.NewProfileResponse(ret)))
}
