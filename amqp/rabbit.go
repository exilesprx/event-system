package amqp

import (
	"fmt"
	"os"
	"strconv"

	"github.com/streadway/amqp"

	"github.com/exilesprx/event-system/log"
)

func Connect() *amqp.Connection {
	user := os.Getenv("AMQP_USER")

	password := os.Getenv("AMQP_PASSWORD")

	host := os.Getenv("AMQP_HOST")

	port, _ := strconv.Atoi(os.Getenv("AMQP_PORT"))

	return connect(user, password, host, port)
}

func connect(user string, password, host string, port int) *amqp.Connection {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port)

	conn, err := amqp.Dial(url)

	log.FailOnError(err, "Failed to connect to RabbitMQ")

	defer closeConnection(conn)

	return conn
}

func channel(connection amqp.Connection) {
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