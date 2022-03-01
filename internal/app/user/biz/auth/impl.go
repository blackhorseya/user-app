package auth

import (
	"strings"
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/user-app/internal/app/user/biz/auth/repo"
	"github.com/blackhorseya/user-app/internal/pkg/entity/er"
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	"github.com/blackhorseya/user-app/internal/pkg/infra/authenticator"
	"github.com/blackhorseya/user-app/internal/pkg/infra/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	auth   authenticator.Authenticator
	jwt    jwt.IJwt
	repo   repo.IRepo
}

// NewImpl return IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo, auth authenticator.Authenticator, jwt jwt.IJwt) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "auth.biz")),
		repo:   repo,
		auth:   auth,
		jwt:    jwt,
	}
}

func (i *impl) GetLoginURL(ctx contextx.Contextx, state string) string {
	return i.auth.AuthCodeURL(ctx, state)
}

func (i *impl) Callback(ctx contextx.Contextx, code string) (info *user.Profile, err error) {
	token, err := i.auth.Exchange(ctx, code)
	if err != nil {
		i.logger.Error(er.ErrExchangeByCode.Error(), zap.Error(err), zap.String("code", code))
		return nil, er.ErrExchangeByCode
	}

	idToken, err := i.auth.VerifyIDToken(ctx, token)
	if err != nil {
		i.logger.Error(er.ErrVerifyIDToken.Error(), zap.Error(err), zap.Any("token", token))
		return nil, er.ErrVerifyIDToken
	}

	claims, err := i.auth.Claims(idToken)
	if err != nil {
		i.logger.Error(er.ErrIDTokenClaims.Error(), zap.Error(err), zap.Any("id_token", idToken))
		return nil, er.ErrIDTokenClaims
	}

	subs := strings.Split(claims["sub"].(string), "|")
	provider := subs[0]
	openID := subs[1]

	exists, err := i.repo.GetUserByOpenID(ctx, provider, openID)
	if err != nil {
		i.logger.Error(er.ErrGetUserByOpenID.Error(), zap.Error(err), zap.String("provider", provider), zap.String("open_id", openID))
		return nil, er.ErrGetUserByOpenID
	}
	if exists != nil {
		exists.Name = claims["name"].(string)
		exists.Nickname = claims["nickname"].(string)
		exists.Email = claims["email"].(string)
		exists.PictureURL = claims["picture"].(string)

		newToken, err := i.jwt.NewToken(exists)
		if err != nil {
			i.logger.Error(er.ErrNewToken.Error(), zap.Error(err), zap.Any("exists", exists))
			return nil, er.ErrNewToken
		}
		exists.Token = newToken

		ret, err := i.repo.UpdateUser(ctx, exists)
		if err != nil {
			i.logger.Error(er.ErrUpdateUser.Error(), zap.Error(err), zap.Any("exists", exists))
			return nil, er.ErrUpdateUser
		}

		return ret, nil
	}

	newUser := &user.Profile{
		ID:           primitive.NewObjectIDFromTimestamp(time.Now()),
		OpenIds:      map[string]string{provider: openID},
		Name:         claims["name"].(string),
		Nickname:     claims["nickname"].(string),
		Email:        claims["email"].(string),
		Token:        "",
		AccessToken:  "",
		RefreshToken: "",
		PictureURL:   claims["picture"].(string),
		CreatedAt:    0,
		UpdatedAt:    0,
	}
	newToken, err := i.jwt.NewToken(newUser)
	if err != nil {
		i.logger.Error(er.ErrNewToken.Error(), zap.Error(err), zap.Any("new_user", newUser))
		return nil, er.ErrNewToken
	}
	newUser.Token = newToken

	ret, err := i.repo.RegisterUser(ctx, newUser)
	if err != nil {
		i.logger.Error(er.ErrRegisterUser.Error(), zap.Error(err), zap.Any("new_user", newUser))
		return nil, er.ErrRegisterUser
	}

	return ret, nil
}

func (i *impl) GetUserByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error) {
	ret, err := i.repo.GetUserByToken(ctx, token)
	if err != nil {
		i.logger.Error(er.ErrGetUserByToken.Error(), zap.Error(err), zap.String("token", token))
		return nil, er.ErrGetUserByToken
	}

	return ret, nil
}

func (i *impl) HasPermission(ctx contextx.Contextx, token, obj, act string) (bool, error) {
	// todo: 2022-03-01|05:26|Sean|impl me
	panic("implement me")
}
