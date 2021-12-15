package system

import (
	"github.com/multipleton/shooter/server/core/handlers/statistics"
	"github.com/multipleton/shooter/server/core/models/player"
	"github.com/multipleton/shooter/server/core/models/system"
)

type Round struct {
	server     *system.Server
	statistics *statistics.RoundStatistics
	players    []*player.Player
}

func NewRound(server *system.Server) *Round {
	statistics := &statistics.RoundStatistics{}
	players := []*player.Player{}
	return &Round{
		server:     server,
		statistics: statistics,
		players:    players,
	}
}
