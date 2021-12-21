package models

import (
	"errors"
	"fmt"

	"github.com/multipleton/shooter/server/core/handlers/statistics"
	"github.com/multipleton/shooter/server/core/models/player"
)

type State struct {
	Statistics *statistics.GameStatistics `json:"statistics"`
	Players    []*player.Player           `json:"players"`
}

func (s *State) FindPlayerByUserId(id int) (*player.Player, error) {
	for _, player := range s.Players {
		if player.User.Id == id {
			return player, nil
		}
	}
	message := fmt.Sprintf("player with id=%d doesn't exists", id)
	return nil, errors.New(message)
}

func NewState(players []*player.Player) *State {
	statistics := &statistics.GameStatistics{}
	return &State{
		Statistics: statistics,
		Players:    players,
	}
}
