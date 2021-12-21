package engine

import (
	"github.com/multipleton/shooter/server/core/models/player"
	"github.com/multipleton/shooter/server/core/models/side"
	"github.com/multipleton/shooter/server/core/models/system"
	"github.com/multipleton/shooter/server/utils"
)

type Game struct {
	server       *system.Server
	loop         *Loop
	state        *State
	eventEmitter *utils.EventEmitter
}

func (g *Game) Start() {
	g.loop.Start(g.Update)
}

func (g *Game) Update() {
	g.eventEmitter.Emit(string(STATE), g.state)
}

func NewGame(
	server *system.Server,
	eventEmitter *utils.EventEmitter,
) *Game {
	loop := NewLoop()
	players := []*player.Player{}
	for i, user := range server.Users {
		player := player.NewPlayer(user, &side.Array[i])
		players = append(players, player)
	}
	state := NewState(players)
	return &Game{
		server:       server,
		loop:         loop,
		state:        state,
		eventEmitter: eventEmitter,
	}
}
