package utils

type EventHandler interface {
	Handle(interface{})
}
