package engine

import (
	"github.com/multipleton/shooter/server/core/handlers/statistics"
	"github.com/multipleton/shooter/server/core/models/player"
)

type State struct {
	statistics *statistics.GameStatistics
	players    []*player.Player
}

func (s *State) Update() {
	// TODO: call update in udp
}

func NewState(players []*player.Player) *State {
	statistics := &statistics.GameStatistics{}
	return &State{
		statistics: statistics,
		players:    players,
	}
}
