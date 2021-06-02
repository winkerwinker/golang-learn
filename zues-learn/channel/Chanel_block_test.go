package chanel

import (
	"fmt"
	"testing"
	"time"
)

// 轻松实现了两者的交互
func TestChanelBlock(t *testing.T) {
	ch1 := make(chan int)
	done := make(chan bool) // 通道
	go func() {
		fmt.Println("子goroutine执行。。。")
		time.Sleep(3 * time.Second)
		data := <-ch1 // 从通道中读取数据
		fmt.Println("data：", data)
		done <- true
	}()
	// 向通道中写数据。。
	time.Sleep(5 * time.Second)
	ch1 <- 100
	// 吐给空
	<-done
	fmt.Println("main。。over")
}
