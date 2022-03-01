package er

import (
	"net/http"

	"github.com/blackhorseya/gocommon/pkg/er"
)

var (
	// ErrUserNotExists means User is NOT exists.
	ErrUserNotExists = er.NewAPPError(http.StatusNotFound, 40410, "User is NOT exists.")
)
