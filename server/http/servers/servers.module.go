package servers

import "github.com/google/wire"

var Module = wire.NewSet(
	NewServersController,
	NewServersService,
	NewServersRepository,
)
