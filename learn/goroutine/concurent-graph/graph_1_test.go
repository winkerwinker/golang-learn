package concurent_graph_test

import "testing"

func Test(t *testing.T) {

	nums := make(chan int)

	go func() {
		// 向通道发送数字42
		nums <- 42
	}()

	// 等待读取出
	<-nums
}
