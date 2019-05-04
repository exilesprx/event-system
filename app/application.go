package app

import (
	"github.com/exilesprx/event-system/amqp"
	"github.com/exilesprx/event-system/log"
	"github.com/joho/godotenv"
)

type Application struct {
	rabbit amqp.Rabbit
}

func (app *Application) Run() {
	loadEnv()

	startAmqpServer(app)
}

func startAmqpServer(app *Application) {
	app.rabbit.Connect()

	defer app.rabbit.Close()

	app.rabbit.Work()
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.FailOnError(err, "Could not load env file")
	}
}