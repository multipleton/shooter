package engine

import (
	"github.com/multipleton/shooter/server/core/models/system"
)

type Manager struct {
	games []*Game
}

func (gc *Manager) FindGameByServerId(serverId int) *Game {
	for _, game := range gc.games {
		if game.server.Id == serverId {
			return game
		}
	}
	return nil // TODO: add error on not found
}

func (gc *Manager) CreateNewGame(server *system.Server) {
	game := NewGame(server)
	gc.games = append(gc.games, game)
}

func NewManager() *Manager {
	return &Manager{games: []*Game{}}
}
