package amqp

import (
	"fmt"
	"github.com/exilesprx/event-system/log"
	"github.com/streadway/amqp"
)

type Rabbit struct {
	connection *amqp.Connection
	queue      amqp.Queue
	channel    *amqp.Channel
}

func (rabbit *Rabbit) Connect(user string, password string, host string, port int) {
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

func (rabbit *Rabbit) Consume() <- chan amqp.Delivery {
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

func (rabbit *Rabbit) Work(channel string, processor MessageProcessor) {
	rabbit.DeclareQueue(channel)

	messages := rabbit.Consume()

	forever := make(chan bool)

	go processor.Process(messages)

	log.Print("Working...")

	<-forever
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
