package models

import (
	"github.com/multipleton/shooter/server/core/models/item"
	"github.com/multipleton/shooter/server/core/models/side"
	"github.com/multipleton/shooter/server/core/models/weapon"
)

type ModelsService struct {
}

func (ms *ModelsService) GetAllWeapons() []weapon.Weapon {
	return weapon.Array
}

func (ms *ModelsService) GetAllItems() []item.Item {
	return item.Array
}

func (ms *ModelsService) GetAllSides() []side.Side {
	return side.Array
}

func NewModelsService() *ModelsService {
	return &ModelsService{}
}
