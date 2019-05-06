package handlers

import (
	"fmt"
	"github.com/exilesprx/event-system/log"
)

const eventName = "test"

type TestHandler struct {
}

func (handler *TestHandler) Handle(message string) {
	msg := fmt.Sprintf("Handled from TestHandler: %s", message)

	log.Print(msg)
}

func (handler *TestHandler) GetEventName() string {
	return eventName
}
