package er

import (
	"net/http"

	"github.com/blackhorseya/gocommon/pkg/er"
)

var (
	// ErrPingDatabase means Ping database is failure.
	ErrPingDatabase = er.NewAPPError(http.StatusInternalServerError, 50001, "Ping database is failure.")
)

var (
	// ErrExchangeByCode means Exchange by code is failure.
	ErrExchangeByCode = er.NewAPPError(http.StatusInternalServerError, 50010, "Exchange by code is failure.")

	// ErrVerifyIDToken means Verify id token is failure.
	ErrVerifyIDToken = er.NewAPPError(http.StatusInternalServerError, 50011, "Verify id token is failure.")

	// ErrIDTokenClaims means ID token claims is failure.
	ErrIDTokenClaims = er.NewAPPError(http.StatusInternalServerError, 50012, "ID token claims is failure.")

	// ErrGetUserByOpenID means Get user by open id is failure.
	ErrGetUserByOpenID = er.NewAPPError(http.StatusInternalServerError, 50013, "Get user by open id is failure.")

	// ErrUpdateUser means Update user profile is failure.
	ErrUpdateUser = er.NewAPPError(http.StatusInternalServerError, 50014, "Update user profile is failure.")

	// ErrNewToken means Create a new token from profile is failure.
	ErrNewToken = er.NewAPPError(http.StatusInternalServerError, 50015, "Create a new token from profile is failure.")

	// ErrRegisterUser means Register user is failure.
	ErrRegisterUser = er.NewAPPError(http.StatusInternalServerError, 50016, "Register user is failure.")

	// ErrGetUserByToken means Get user by token is failure.
	ErrGetUserByToken = er.NewAPPError(http.StatusInternalServerError, 50017, "Get user by token is failure.")
)
