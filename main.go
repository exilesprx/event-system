package main

import (
	"github.com/exilesprx/event-system/amqp"

	"strconv"

	"os"

	"github.com/exilesprx/event-system/log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	user := os.Getenv("AMQP_USER")

	password := os.Getenv("AMQP_PASSWORD")

	host := os.Getenv("AMQP_HOST")

	port, _ := strconv.Atoi(os.Getenv("AMQP_PORT"))

	amqp.Connect(user, password, host, port)
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.FailOnError(err, "Could not load env file")
	}
}