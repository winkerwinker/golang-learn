package chanel

import (
	"fmt"
	"testing"
)

// 死锁
func TestChanelDeadLock(t *testing.T) {

	//all goroutines are asleep - deadlock!
	ch := make(chan int)
	go func() {
		<-	ch
	}()

	ch <- 5

	fmt.Printf("over")
}

