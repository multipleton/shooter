package udp

type UDPResultHanlder struct {
	callback func(interface{})
}

func (udprh *UDPResultHanlder) Handle(object interface{}) {
	udprh.callback(object)
}

func NewUDPResultHandler(callback func(interface{})) *UDPResultHanlder {
	return &UDPResultHanlder{
		callback: callback,
	}
}
