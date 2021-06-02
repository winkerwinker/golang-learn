package main

import (
	"fmt"
	"testing"
	"time"
)



func TestArrowUse(t *testing.T) {
	v := "123"
	ch := make(chan string)
	ch <- v    // 发送值v到Channel ch中

	time.Sleep(time.Duration(2)*time.Second)

	fmt.Print(ch)



	// v := <-ch  // 从Channel ch中接收数据，并将数据赋值给v


}
