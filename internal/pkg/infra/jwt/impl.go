package jwt

import (
	"time"

	"github.com/blackhorseya/user-app/internal/pkg/entity/er"
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Options declare jwt configuration
type Options struct {
	Project   string
	Signature string
}

// NewOptions serve caller to create Options
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, err
	}

	return o, nil
}

type impl struct {
	o      *Options
	logger *zap.Logger
}

// NewImpl return IJwt
func NewImpl(o *Options, logger *zap.Logger) (IJwt, error) {
	return &impl{
		o:      o,
		logger: logger,
	}, nil
}

func (i *impl) NewToken(info *user.Profile) (string, error) {
	claims := TokenClaims{
		ID:      info.ID.Hex(),
		Name:    info.Name,
		Picture: info.PictureURL,
		Email:   info.Email,
		StandardClaims: jwt.StandardClaims{
			Issuer: i.o.Project,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(i.o.Signature))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (i *impl) ValidateToken(signedToken string) (claims *TokenClaims, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(i.o.Signature), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, er.ErrInvalidToken
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, er.ErrExpiredToken
	}

	return claims, nil
}
