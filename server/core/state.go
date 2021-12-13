package core

import "github.com/multipleton/shooter/server/core/models/system"

type State struct {
	users   []*system.User
	servers []*system.Server
}

func NewState() *State {
	return &State{
		users:   []*system.User{},
		servers: []*system.Server{},
	}
}
