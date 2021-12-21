package engine

import (
	"log"

	"github.com/multipleton/shooter/server/core/models/system"
	"github.com/multipleton/shooter/server/utils"
)

type Manager struct {
	games        []*Game
	eventEmitter *utils.EventEmitter
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
	game := NewGame(server, gc.eventEmitter)
	gc.games = append(gc.games, game)
	log.Printf("game for server %d created", server.Id)
	game.Start()
}

func NewManager(eventEmitter *utils.EventEmitter) *Manager {
	return &Manager{
		games:        []*Game{},
		eventEmitter: eventEmitter,
	}
}
