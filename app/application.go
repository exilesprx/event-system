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
	rabbit amqp.Rabbit
}

func (app *Application) Run() {
	loadEnv()

	app.startAmqpServer()
}

func (app *Application) Shutdown() {
	defer app.rabbit.Close()
}

func (app *Application) startAmqpServer() {
	user := os.Getenv("AMQP_USER")

	password := os.Getenv("AMQP_PASSWORD")

	host := os.Getenv("AMQP_HOST")

	port, _ := strconv.Atoi(os.Getenv("AMQP_PORT"))

	rabbit := amqp.Rabbit{}

	rabbit.Connect(user, password, host, port)

	processor := setupEventProcessor()

	rabbit.Work(os.Getenv("AMQP_CHANNEL"), processor)

	app.rabbit = rabbit
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
