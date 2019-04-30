package amqp

import (
	"fmt"
	"os"
	"strconv"

	"github.com/streadway/amqp"

	"github.com/exilesprx/event-system/log"
)

type Rabbit struct {
	connection *amqp.Connection
	queue amqp.Queue
	channel *amqp.Channel
}

func (rabbit *Rabbit) Connect() {
	user := os.Getenv("AMQP_USER")

	password := os.Getenv("AMQP_PASSWORD")

	host := os.Getenv("AMQP_HOST")

	port, _ := strconv.Atoi(os.Getenv("AMQP_PORT"))

	rabbit.connection = connect(user, password, host, port)
}

func (rabbit *Rabbit) DeclareQueue(name string) {
	rabbit.channel = channel(rabbit.connection)

	q, err := rabbit.channel.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil)

	if err != nil {
		log.FailOnError(err, "Failed to declare queue")
	}

	rabbit.queue = q
}

func (rabbit *Rabbit) Consume() <-chan amqp.Delivery {
	messages, _ := rabbit.channel.Consume(
		rabbit.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)

	return messages
}

func connect(user string, password, host string, port int) *amqp.Connection {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port)

	conn, err := amqp.Dial(url)

	log.FailOnError(err, "Failed to connect to RabbitMQ")

	defer closeConnection(conn)

	return conn
}

func channel(connection *amqp.Connection) *amqp.Channel {
	channel, err := connection.Channel()

	log.FailOnError(err, "Failed to create channel")

	defer closeChannel(channel)

	return channel
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
