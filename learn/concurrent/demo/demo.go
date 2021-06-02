package concurrent

import (
	"fmt"
	"time"
)

func Demo1(s string ) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("  -> %d %s \n ", i, s)

	}

}

func Run() {
	go Demo1("hello")
	Demo1("world")
}
