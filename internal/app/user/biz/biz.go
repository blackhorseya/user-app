package biz

import (
	"github.com/blackhorseya/user-app/internal/app/user/biz/health"
	"github.com/google/wire"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
)
