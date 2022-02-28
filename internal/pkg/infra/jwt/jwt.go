package jwt

import (
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	"github.com/golang-jwt/jwt"
	"github.com/google/wire"
)

// TokenClaims declare custom claims
type TokenClaims struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Email   string `json:"email"`
	jwt.StandardClaims
}

// IJwt declare a jwt factory functions
//go:generate mockery --name=IJwt
type IJwt interface {
	NewToken(info *user.Profile) (string, error)

	ValidateToken(signedToken string) (claims *TokenClaims, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewOptions, NewImpl)
