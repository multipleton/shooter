package handlers

import (
	"encoding/json"
	"log"

	"github.com/multipleton/shooter/server/core/models/player"
	"github.com/multipleton/shooter/server/utils"
)

type MovementHandler struct {
}

func (mh *MovementHandler) Hanlde(jsonObject utils.JSONObject) {
	var position player.Position
	err := json.Unmarshal([]byte(jsonObject.Value), &position)
	if err != nil {
		log.Printf("cannot parse player position from: %s", jsonObject)
		return
	}
	// TODO: update user position
}
