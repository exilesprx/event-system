package events

type Handler interface {
	Handle(message string)
	GetEventName() string
}

type EventHandler struct {
	eventHandlers map[string]Handler
}

func New() EventHandler {
	return EventHandler{
		eventHandlers: make(map[string]Handler),
	}
}

func (eventHandler *EventHandler) GetEventHandlers() map[string]Handler {
	return eventHandler.eventHandlers
}

func (eventHandler *EventHandler) RegisterHandler(handler Handler) {
	eventHandler.eventHandlers[handler.GetEventName()] = handler
}