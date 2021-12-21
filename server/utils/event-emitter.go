package utils

import (
	"errors"
	"fmt"
)

type EventEmitter struct {
	events map[string]*[]*EventHandler
}

func (el *EventEmitter) On(event string, handler *EventHandler) {
	handlers := el.events[event]
	if handlers == nil {
		handlers = &[]*EventHandler{}
	}
	(*handlers) = append((*handlers), handler)
}

func (el *EventEmitter) Emit(event string, data interface{}) error {
	handlers := el.events[event]
	if handlers == nil {
		message := fmt.Sprintf("event %s doesn't exists", event)
		return errors.New(message)
	}
	for _, handler := range *handlers {
		(*handler).Handle(data)
	}
	return nil
}

func NewEventEmitter() *EventEmitter {
	events := make(map[string]*[]*EventHandler)
	return &EventEmitter{
		events: events,
	}
}
