package amqp

import (
	"fmt"

	"github.com/streadway/amqp"

	"github.com/exilesprx/event-system/log"
)

func Connect(host string, port int) *amqp.Connection {
	url := fmt.Sprintf("amqp://guest:guest@%s:%d", host, port)

	conn, err := amqp.Dial(url)

	log.FailOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	return conn
}