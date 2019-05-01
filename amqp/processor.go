package amqp

import "github.com/streadway/amqp"

type MessageProcessor struct {
}

func (processor *MessageProcessor) Process(<-chan amqp.Delivery) {
	// TODO: Add logic to process message, ex: fire off event
}
