package main

import (
	"github.com/exilesprx/event-system/amqp"

	"github.com/exilesprx/event-system/log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	amqp.Connect()
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.FailOnError(err, "Could not load env file")
	}
}