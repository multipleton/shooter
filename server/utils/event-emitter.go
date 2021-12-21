package utils

import (
	"log"
)

type EventEmitter struct {
	events map[string]*[]EventHandler
}

func (el *EventEmitter) On(event string, handler EventHandler) {
	if el.events[event] == nil {
		el.events[event] = &[]EventHandler{}
	}
	*el.events[event] = append((*el.events[event]), handler)
}

func (el *EventEmitter) Emit(event string, data interface{}) {
	handlers := el.events[event]
	if handlers == nil {
		log.Printf("event %s doesn't exists", event)
		return
	}
	for _, handler := range *handlers {
		handler.Handle(data)
	}
}

func NewEventEmitter() *EventEmitter {
	events := make(map[string]*[]EventHandler)
	return &EventEmitter{
		events: events,
	}
}
