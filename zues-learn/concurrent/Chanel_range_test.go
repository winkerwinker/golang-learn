package chanel

import (
	"fmt"
	"testing"
	"time"
)

// 使用range处理chanel
func TestChanel3(t *testing.T) {
	// todo 这里会用作计时吗
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}
