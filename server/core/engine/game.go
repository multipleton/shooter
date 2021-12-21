package engine

import (
	"github.com/multipleton/shooter/server/core/engine/event"
	"github.com/multipleton/shooter/server/core/handlers"
	"github.com/multipleton/shooter/server/core/models"
	"github.com/multipleton/shooter/server/core/models/player"
	"github.com/multipleton/shooter/server/core/models/side"
	"github.com/multipleton/shooter/server/core/models/system"
	"github.com/multipleton/shooter/server/utils"
)

type Game struct {
	server        *system.Server
	loop          *Loop
	state         *models.State
	eventEmitter  *utils.EventEmitter
	eventHandlers []utils.EventHandler
}

func (g *Game) Init() {
	movementHandler := handlers.NewMovementHandler(g.state)
	bulletHandler := handlers.NewBulletHandler(g.state)
	itemHandler := handlers.NewItemHandler(g.state)
	g.eventEmitter.On(string(event.POSITION), movementHandler)
	g.eventEmitter.On(string(event.BULLET), bulletHandler)
	g.eventEmitter.On(string(event.POSITION), itemHandler)
	g.eventHandlers = []utils.EventHandler{
		movementHandler,
		bulletHandler,
		itemHandler,
	}
}

func (g *Game) Start() {
	g.loop.Start(g.Update)
}

func (g *Game) Update() {
	g.eventEmitter.Emit(string(event.STATE), g.state)
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
	state := models.NewState(players)
	game := &Game{
		server:       server,
		loop:         loop,
		state:        state,
		eventEmitter: eventEmitter,
	}
	game.Init()
	return game
}
