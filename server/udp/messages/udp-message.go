package messages

import "github.com/multipleton/shooter/server/utils"

type UDPMessage struct {
	MessageType UDPMessageType   `json:"type"`
	Data        utils.JSONObject `json:"data"`
}
