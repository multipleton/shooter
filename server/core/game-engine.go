package core

type GameEngine struct {
	State *State
}

func NewGameEngine() *GameEngine {
	return &GameEngine{} // TODO: init state
}
