package core

type GameEngine struct {
	state *State
}

func NewGameEngine() *GameEngine {
	state := NewState()
	return &GameEngine{state: state} // TODO: init state
}
