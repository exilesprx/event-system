package app

import (
	"github.com/exilesprx/event-system/amqp"
	"github.com/exilesprx/event-system/events"
	"github.com/exilesprx/event-system/handlers"
	"github.com/exilesprx/event-system/log"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Application struct {
}

func (app *Application) Run() {
	loadEnv()

	rabbit := startAmqpServer()

	defer rabbit.Close()

	rabbit.Work(os.Getenv("AMQP_CHANNEL"), setupEventProcessor())
}

func startAmqpServer() amqp.Rabbit {
	user := os.Getenv("AMQP_USER")

	password := os.Getenv("AMQP_PASSWORD")

	host := os.Getenv("AMQP_HOST")

	port, _ := strconv.Atoi(os.Getenv("AMQP_PORT"))

	rabbit := amqp.Rabbit{}

	rabbit.Connect(user, password, host, port)

	return rabbit
}

func setupEventProcessor() amqp.MessageProcessor {
	processor := amqp.NewMessageProcessor()

	handler := events.New()

	handler.RegisterHandler(&handlers.TestHandler{})

	processor.RegisterHandlers(handler.GetEventHandlers())

	return processor
}

func loadEnv() {
	err := godotenv.Load()

	log.FailOnError(err, "Could not load env file")
}