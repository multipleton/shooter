package system

import (
	"github.com/multipleton/shooter/server/core/handlers/statistics"
	"github.com/multipleton/shooter/server/core/models/player"
	"github.com/multipleton/shooter/server/core/models/system"
)

type Round struct {
	Server     *system.Server
	Statistics *statistics.RoundStatistics
	Players    []*player.Player
}
