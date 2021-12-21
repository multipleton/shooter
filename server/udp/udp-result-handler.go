package udp

type UDPResultHandler struct {
	callback func(interface{})
}

func (udprh *UDPResultHandler) Handle(object interface{}) {
	udprh.callback(object)
}

func NewUDPResultHandler(callback func(interface{})) *UDPResultHandler {
	return &UDPResultHandler{
		callback: callback,
	}
}
