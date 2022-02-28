package er

import (
	"net/http"

	"github.com/blackhorseya/gocommon/pkg/er"
)

var (
	// ErrPingDatabase means Ping database is failure.
	ErrPingDatabase = er.NewAPPError(http.StatusInternalServerError, 50001, "Ping database is failure.")
)
