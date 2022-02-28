package er

import (
	"net/http"

	"github.com/blackhorseya/gocommon/pkg/er"
)

var (
	// ErrMissingToken means missing jwt in header
	ErrMissingToken = er.NewAPPError(http.StatusUnauthorized, 40100, "missing jwt token")

	// ErrAuthHeaderFormat means must provide Authorization header with format `Bearer {jwt}`
	ErrAuthHeaderFormat = er.NewAPPError(http.StatusUnauthorized, 40101, "Must provide Authorization header with format `Bearer {jwt}`")

	// ErrExpiredToken means jwt is expired
	ErrExpiredToken = er.NewAPPError(http.StatusUnauthorized, 40102, "jwt is expired")

	// ErrInvalidToken means jwt is invalid
	ErrInvalidToken = er.NewAPPError(http.StatusUnauthorized, 40103, "jwt is invalid")
)
