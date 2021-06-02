package chanel

import (
	"fmt"
	"testing"
	"time"
)


//用于超时处理
func TestChanel5(t *testing.T) {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
		// todo 这里是什么意思，再等一秒的意思吗
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}

