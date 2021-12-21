package player

import (
	"github.com/multipleton/shooter/server/core/models/side"
	"github.com/multipleton/shooter/server/core/models/system"
	"github.com/multipleton/shooter/server/core/models/weapon"
)

type Player struct {
	User   *system.User
	Side   *side.Side
	HP     uint16
	Armor  uint16
	Weapon weapon.Weapon
}

func NewPlayer(
	user *system.User,
	side *side.Side,
) *Player {
	return &Player{
		User:   user,
		Side:   side,
		HP:     100,
		Armor:  0,
		Weapon: weapon.KNIFE,
	}
}
