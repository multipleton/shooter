package messages

const (
	CONNECT    UDPMessageType = "connect"
	DISCONNECT UDPMessageType = "disconnect"
	STATE      UDPMessageType = "state"
	POSITION   UDPMessageType = "position"
	BULLET     UDPMessageType = "bullet"
)

var Array = []UDPMessageType{
	CONNECT,
	DISCONNECT,
	STATE,
	POSITION,
}
