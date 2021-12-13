package models

import (
	"github.com/multipleton/shooter/server/core/models/item"
	"github.com/multipleton/shooter/server/core/models/side"
	"github.com/multipleton/shooter/server/core/models/weapon"
)

type ModelsService struct {
}

func (ms *ModelsService) GetAllWeapons() []weapon.Weapon {
	return []weapon.Weapon{
		weapon.KNIFE,
		weapon.FIVESEVEN,
		weapon.DEAGLE,
		weapon.MP5,
		weapon.P90,
		weapon.M4A1,
		weapon.AK47,
		weapon.XM1014,
		weapon.VT,
		weapon.M60,
		weapon.AWP,
	}
}

func (ms *ModelsService) GetAllItems() []item.Item {
	return []item.Item{
		item.HP,
		item.ARMOR,
		item.DAMAGE,
		item.SPEED,
		item.CAKE,
	}
}

func (ms *ModelsService) GetAllSides() []side.Side {
	return []side.Side{
		side.RED,
		side.GREEN,
		side.BLUE,
		side.YELLOW,
	}
}

func NewModelsService() *ModelsService {
	return &ModelsService{}
}
