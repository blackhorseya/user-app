// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/gocommon/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	oauth2 "golang.org/x/oauth2"

	oidc "github.com/coreos/go-oidc/v3/oidc"
)

// Authenticator is an autogenerated mock type for the Authenticator type
type Authenticator struct {
	mock.Mock
}

// AuthCodeURL provides a mock function with given fields: ctx, state, opts
func (_m *Authenticator) AuthCodeURL(ctx contextx.Contextx, state string, opts ...oauth2.AuthCodeOption) string {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, state)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, ...oauth2.AuthCodeOption) string); ok {
		r0 = rf(ctx, state, opts...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Claims provides a mock function with given fields: idToken
func (_m *Authenticator) Claims(idToken *oidc.IDToken) (map[string]interface{}, error) {
	ret := _m.Called(idToken)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(*oidc.IDToken) map[string]interface{}); ok {
		r0 = rf(idToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*oidc.IDToken) error); ok {
		r1 = rf(idToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exchange provides a mock function with given fields: ctx, code, opts
func (_m *Authenticator) Exchange(ctx contextx.Contextx, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, code)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *oauth2.Token
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, ...oauth2.AuthCodeOption) *oauth2.Token); ok {
		r0 = rf(ctx, code, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string, ...oauth2.AuthCodeOption) error); ok {
		r1 = rf(ctx, code, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyIDToken provides a mock function with given fields: ctx, token
func (_m *Authenticator) VerifyIDToken(ctx contextx.Contextx, token *oauth2.Token) (*oidc.IDToken, error) {
	ret := _m.Called(ctx, token)

	var r0 *oidc.IDToken
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *oauth2.Token) *oidc.IDToken); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oidc.IDToken)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *oauth2.Token) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}