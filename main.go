package main

import (
	"github.com/exilesprx/event-system/amqp"
	"os"

	"github.com/exilesprx/event-system/log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	rabbit := amqp.Rabbit{}

	rabbit.Connect()

	rabbit.DeclareQueue(os.Getenv("AMQP_CHANNEL"))

	rabbit.Consume()
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.FailOnError(err, "Could not load env file")
	}
}