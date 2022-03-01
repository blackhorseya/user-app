package er

import (
	"net/http"

	"github.com/blackhorseya/gocommon/pkg/er"
)

var (
	// ErrInvalidState means Invalid state parameter
	ErrInvalidState = er.NewAPPError(http.StatusBadRequest, 40010, "Invalid state parameter.")
)
