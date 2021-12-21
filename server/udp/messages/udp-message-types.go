package messages

const (
	CONNECT    UDPMessageType = "connect"
	DISCONNECT UDPMessageType = "disconnect"
)

var Array = []UDPMessageType{
	CONNECT,
	DISCONNECT,
}
