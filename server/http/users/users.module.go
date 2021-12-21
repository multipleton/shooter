package users

import "github.com/google/wire"

var Module = wire.NewSet(
	NewUsersService,
	NewUsersRepository,
)
