package messages

const (
	CONNECT    UDPMessageType = "connect"
	DISCONNECT UDPMessageType = "disconnect"
	STATE      UDPMessageType = "state"
)

var Array = []UDPMessageType{
	CONNECT,
	DISCONNECT,
	STATE,
}
