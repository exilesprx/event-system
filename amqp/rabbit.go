package amqp

import (
	"fmt"

	"github.com/streadway/amqp"

	"github.com/exilesprx/event-system/log"
)

func Connect(user string, password, host string, port int) *amqp.Connection {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port)

	conn, err := amqp.Dial(url)

	log.FailOnError(err, "Failed to connect to RabbitMQ")

	defer closeConnection(conn)

	return conn
}

func Channel(connection amqp.Connection) {
	channel, err := connection.Channel()

	log.FailOnError(err, "Failed to create channel")

	defer closeChannel(channel)
}

func closeConnection(connection *amqp.Connection) {
	err := connection.Close()

	if err != nil {
		log.FailOnError(err, "Failed to close connection")
	}
}

func closeChannel(channel *amqp.Channel) {
	err := channel.Close()

	if err != nil {
		log.FailOnError(err, "Failed to close connection")
	}
}