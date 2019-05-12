package events

import (
	"fmt"
	"testing"
)

type TestHandler struct {
	EventName string
}

func (handler *TestHandler) Handle(message string) {

}

func (handler *TestHandler) GetEventName() string {
	return handler.EventName
}

func TestRegisterHandler(t *testing.T) {
	t.Run("1", testRegisterHandler(1))

	t.Run("5", testRegisterHandler(5))

	t.Run("20", testRegisterHandler(20))
}

func testRegisterHandler(eventCount int) func(t *testing.T) {
	return func(t *testing.T) {
		events := New()
		handler := &TestHandler{}

		for i := 0; i < eventCount; i++ {
			handler.EventName = fmt.Sprintf("test%d", i)

			events.RegisterHandler(handler)
		}

		allEvents := events.GetEventHandlers()

		if len(allEvents) != eventCount {
			t.Fail()
		}
	}
}