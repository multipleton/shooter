package engine

import (
	"github.com/multipleton/shooter/server/core/handlers/statistics"
	"github.com/multipleton/shooter/server/core/models/player"
)

type State struct {
	Statistics *statistics.GameStatistics `json:"statistics"`
	Players    []*player.Player           `json:"players"`
}

func (s *State) Update() {
	// TODO: call update in udp
}

func NewState(players []*player.Player) *State {
	statistics := &statistics.GameStatistics{}
	return &State{
		Statistics: statistics,
		Players:    players,
	}
}
