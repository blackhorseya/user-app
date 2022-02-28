package authenticator

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/google/wire"
	"golang.org/x/oauth2"
)

// Authenticator declare authenticator functions
//go:generate mockery --name=Authenticator
type Authenticator interface {
	AuthCodeURL(ctx contextx.Contextx, state string, opts ...oauth2.AuthCodeOption) string

	Exchange(ctx contextx.Contextx, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)

	VerifyIDToken(ctx contextx.Contextx, token *oauth2.Token) (*oidc.IDToken, error)

	Claims(idToken *oidc.IDToken) (profile map[string]interface{}, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewOptions, NewImpl)
