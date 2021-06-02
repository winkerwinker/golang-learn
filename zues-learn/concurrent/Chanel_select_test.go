package chanel

import (
	"fmt"
	"testing"
)

func TestChanel4(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(<-c)
		}
		// 使用quit 来协调两边的顺序
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// 写进c
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
