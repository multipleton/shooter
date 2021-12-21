package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/multipleton/shooter/server/core/models"
	"github.com/multipleton/shooter/server/core/models/player"
	"github.com/multipleton/shooter/server/utils"
)

type ItemHandler struct {
	state *models.State
}

func (ih *ItemHandler) Handle(object interface{}) {
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
	// TODO: check position with items position
	fmt.Println("item", id, position)
}

func NewItemHandler(state *models.State) *ItemHandler {
	return &ItemHandler{state: state}
}
