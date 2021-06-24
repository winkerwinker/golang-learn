package concurent_graph_test

import "testing"

func Test_2(t *testing.T) {

	nums := make(chan int)

	for i := 0; i < 10; i++ {
		go func(num int) {
			// 向通道发送数字42
			nums <- num
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-nums
	}
	// 等待读取出

}
