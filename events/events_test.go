package events

import (
	"testing"
)

type TestHandler struct {
}

func (handler *TestHandler) Handle(message string) {

}

func (handler *TestHandler) GetEventName() string {
	return "test"
}

func TestRegisterHandler(t *testing.T) {
	events := New()
	handler := &TestHandler{}

	events.RegisterHandler(handler)

	allEvents := events.GetEventHandlers()

	if len(allEvents) != 1 {
		t.Fail()
	}
}