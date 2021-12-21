package player

import (
	"github.com/multipleton/shooter/server/core/models/side"
	"github.com/multipleton/shooter/server/core/models/system"
	"github.com/multipleton/shooter/server/core/models/weapon"
)

type Player struct {
	User     *system.User
	Side     *side.Side
	HP       uint16
	Armor    uint16
	Weapon   weapon.Weapon
	Position Position
}

func NewPlayer(
	user *system.User,
	side *side.Side,
) *Player {
	position := Position{
		X:     0,
		Y:     0,
		Angle: 90,
	}
	return &Player{
		User:     user,
		Side:     side,
		HP:       100,
		Armor:    0,
		Weapon:   weapon.KNIFE,
		Position: position,
	}
}
