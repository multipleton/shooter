package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/multipleton/shooter/server/core/models"
	"github.com/multipleton/shooter/server/core/models/player"
	"github.com/multipleton/shooter/server/utils"
)

type BulletHandler struct {
	state *models.State
}

func (bh *BulletHandler) Handle(object interface{}) {
	args, casted := object.([]interface{})
	if !casted {
		log.Printf("invalid input: %s", object)
		return
	}
	id, casted := args[0].(int)
	if !casted {
		log.Printf("could not cast id from: %s", args)
		return
	}
	jsono, casted := args[1].(utils.JSONObject)
	if !casted {
		log.Printf("could not cast JSONObject from: %s", args)
		return
	}
	var position player.Position
	err := json.Unmarshal([]byte(jsono.Value), &position)
	if err != nil {
		log.Printf("invalid player position: %s", object)
		return
	}
	// TODO: run separate goroutine to calculate bullet position and collisions
	fmt.Println("bullet", id, position)
}

func NewBulletHandler(state *models.State) *BulletHandler {
	return &BulletHandler{state: state}
}
