package authenticator

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

// Options declare authenticator configuration
type Options struct {
	Domain       string
	ClientID     string
	ClientSecret string
	CallbackURL  string
	Endpoint     string
}

// NewOptions return *Options
func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	o := new(Options)

	err := v.UnmarshalKey("authenticator", o)
	if err != nil {
		return nil, err
	}

	logger.Info("load authenticator options success")

	return o, nil
}

type impl struct {
	provider *oidc.Provider
	conf     oauth2.Config
}

// NewImpl return Authenticator
func NewImpl(o *Options) (Authenticator, error) {
	provider, err := oidc.NewProvider(contextx.Background(), "https://"+o.Domain+"/")
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     o.ClientID,
		ClientSecret: o.ClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  o.CallbackURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &impl{
		provider: provider,
		conf:     conf,
	}, nil
}

func (i *impl) AuthCodeURL(ctx contextx.Contextx, state string, opts ...oauth2.AuthCodeOption) string {
	return i.conf.AuthCodeURL(state, opts...)
}

func (i *impl) Exchange(ctx contextx.Contextx, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return i.conf.Exchange(ctx, code, opts...)
}

func (i *impl) VerifyIDToken(ctx contextx.Contextx, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: i.conf.ClientID,
	}

	return i.provider.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func (i *impl) Claims(idToken *oidc.IDToken) (profile map[string]interface{}, err error) {
	var ret map[string]interface{}
	err = idToken.Claims(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
