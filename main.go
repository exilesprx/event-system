package main

import (
	"github.com/exilesprx/event-system/amqp"
)

const host = "localhost"
const port = 5672

func main() {
	amqp.Connect(host, port)
}