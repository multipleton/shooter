package player

import (
	"github.com/multipleton/shooter/server/core/models/side"
	"github.com/multipleton/shooter/server/core/models/system"
	"github.com/multipleton/shooter/server/core/models/weapon"
)

type Player struct {
	User   *system.User
	Side   *side.Side
	HP     HP
	Armor  Armor
	Weapon weapon.Weapon
}
