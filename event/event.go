package event

import (
	"reflect"
	"sync"
)

type Event interface {
}

type Listener interface {
	handler(event Event)
	register()
}

func Register(listener Listener, event Event) {
	eventManager.registerMap[listener] = event
}

func Post(event Event) {
	go func() { eventManager.eventQueue <- event }()
}

var eventManager *EventManager
var mutex sync.Mutex

func init() {
	eventManager = &EventManager{registerMap: make(map[Listener]Event, 100), eventQueue: make(chan Event, 10240)}
	go handler()
}

func handler() {
	select {
	case event := <-eventManager.eventQueue:
		dispatch(event)
	}
}

func dispatch(event Event) {
	mutex.Lock()
	defer mutex.Unlock()

	listenerSlice := make([]Listener, 0, 10)
	for k, v := range eventManager.registerMap {
		index := 0
		if reflect.TypeOf(event) == reflect.TypeOf(v) {
			listenerSlice[index] = k
			index++
		}
	}

	for _, v := range listenerSlice {
		go func() {
			v.handler(event)
		}()
	}
}

// 事件监听器
type EventManager struct {
	registerMap map[Listener]Event
	eventQueue  chan Event
}
