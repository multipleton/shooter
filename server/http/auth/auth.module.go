package auth

import "github.com/google/wire"

var Module = wire.NewSet(
	NewAuthController,
	NewAuthService,
)
