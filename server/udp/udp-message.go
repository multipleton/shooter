package udp

import (
	"github.com/multipleton/shooter/server/utils"
)

type UDPMessage struct {
	MessageType string           `json:"type"`
	Data        utils.JSONObject `json:"data"`
}
