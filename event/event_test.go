package event

import (
	"fmt"
	"testing"
	"time"
)

type TestEvent struct {
	name string
	data string
	time time.Time
}

func (event *TestEvent) Name() string {
	return event.name
}

type TestListener struct {
	name string
	event Event
}

func (listener *TestListener)Name() string  {
	return listener.name
}

func (listener *TestListener)Handler(event Event)  {
	fmt.Println(event.Name())
}

func (listener *TestListener)Register()  {
	Register(listener, listener.event)
}


func TestPost(t *testing.T) {
	event := &TestEvent{name: "testEvent", data: "test data", time: time.Now()}


	Post(event)
}
