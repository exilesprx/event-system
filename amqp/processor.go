package amqp

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/exilesprx/event-system/events"
	"github.com/exilesprx/event-system/log"
	"github.com/streadway/amqp"
)

type MessageProcessor struct {
	bus EventBus.Bus
}

func NewMessageProcessor() MessageProcessor {
	processor := MessageProcessor{
		bus: EventBus.New(),
	}

	return processor
}

func (processor *MessageProcessor) RegisterHandlers(handlers map[string]events.Handler) {
	for topic, handler := range handlers {
		err := processor.bus.Subscribe(topic, handler.Handle)

		log.FailOnError(err, "An error occurred")
	}
}

func (processor *MessageProcessor) Process(messages <-chan amqp.Delivery) {
	for message := range messages {
		processor.processMessage(message)
	}
}

func (processor *MessageProcessor) processMessage(message amqp.Delivery) {
	msg := fmt.Sprintf("%s", message.Body)

	processor.bus.Publish("test", msg)
}