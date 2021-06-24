package event_bus_test

import (
	"fmt"
	"runtime"
	"testing"
)

// how to get the func name in a func
func Test(t *testing.T) {
	Foo()
}

func Foo() {
	fmt.Printf("我是 %s, 谁在调用我?\n", printMyName())
	Bar()
}
func Bar() {
	fmt.Printf("我是 %s, 谁又在调用我?\n", printMyName())
}

// todo 对自身对调用栈对一种掌控
func printMyName() string {
	//Up two level
	// Caller reports(举报) file
	caller, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(caller).Name()
}
