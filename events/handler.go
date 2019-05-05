package events

import (
	"github.com/exilesprx/event-system/handlers"
)

type Handler interface {
	Handle(message string)
}

type EventHandler struct {
	topics map[string]Handler
}

func (handler *EventHandler) Setup() {
	handler.topics = make(map[string]Handler)

	handler.topics["test"] = &handlers.TestHandler{}
}

func (handler *EventHandler) GetTopics() map[string]Handler {
	return handler.topics
}