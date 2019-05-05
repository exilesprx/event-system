package amqp

import (
	"fmt"
	"github.com/streadway/amqp"
)

type MessageProcessor struct {
}

func (processor *MessageProcessor) Process(messages <-chan amqp.Delivery) {

	for message := range messages {
		processMessage(message)
	}
}

func processMessage(message amqp.Delivery) {
	fmt.Printf("Event: %s", message.Body)
}