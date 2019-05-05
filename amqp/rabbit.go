package amqp

import (
	"fmt"
	"github.com/exilesprx/event-system/log"
	"github.com/streadway/amqp"
	"os"
	"strconv"
)

type Rabbit struct {
	connection *amqp.Connection
	queue      amqp.Queue
	channel    *amqp.Channel
	processor  MessageProcessor
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

	log.FailOnError(err, "Failed to declare queue")

	rabbit.queue = q
}

func (rabbit *Rabbit) Consume() chan bool {
	messages, _ := rabbit.channel.Consume(
		rabbit.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)

	forever := make(chan bool)

	go rabbit.processor.Process(messages)

	return forever
}

func (rabbit *Rabbit) Work() {
	rabbit.Connect()

	rabbit.DeclareQueue(os.Getenv("AMQP_CHANNEL"))

	rabbit.processor.Setup()

	process := rabbit.Consume()

	log.Print("Working...")

	<-process
}

func (rabbit *Rabbit) Close() {
	closeChannel(rabbit.channel)

	closeConnection(rabbit.connection)
}

func connect(user string, password, host string, port int) *amqp.Connection {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port)

	conn, err := amqp.Dial(url)

	log.FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

func channel(connection *amqp.Connection) *amqp.Channel {
	channel, err := connection.Channel()

	log.FailOnError(err, "Failed to create channel")

	return channel
}

func closeConnection(connection *amqp.Connection) {
	err := connection.Close()

	log.FailOnError(err, "Failed to close connection")
}

func closeChannel(channel *amqp.Channel) {
	err := channel.Close()

	log.FailOnError(err, "Failed to close channel")
}
