package models

import "github.com/google/wire"

var Module = wire.NewSet(NewModelsService, NewModelsController)
