package eventbus

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

func (bus AsyncEventBus) Subscribe(topic string, f interface{}) error {

	bus.lock.Lock()
	defer bus.lock.Unlock()
	handler := reflect.ValueOf(f)

	if handler.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a func")
	}

	values, ok := bus.handlers[topic]
	if !ok {
		values = []reflect.Value{}
	}

	//todo reflect.Value 到底是啥ß
	values = append(values, handler)
	bus.handlers[topic] = values

	return nil
}

func (bus AsyncEventBus) Publish(topic string, args ...interface{}) {
	//panic("implement me")
	handlers, ok := bus.handlers[topic]

	if !ok {
		fmt.Println("not found handlers in topic:", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i := range handlers {
		// func (v Value) Call(in []Value) []Value {
		go handlers[i].Call(params)
	}

}

func sub1(msg1, msg2 string) {
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("sub1, %s %s\n", msg1, msg2)
}

func sub2(msg1, msg2 string) {
	fmt.Printf("sub2, %s %s\n", msg1, msg2)
}

func main() {

	bus := NewAsyncEventBus()
	bus.Subscribe("topic:1", sub1)
	bus.Subscribe("topic:1", sub2)
	// 使用
	bus.Publish("topic:1", "test1", "test2")
	bus.Publish("topic:1", "testA", "testB")
	time.Sleep(1 * time.Second)

}
