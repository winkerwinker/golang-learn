package concurent_graph_test

import (
	"fmt"
	"testing"
	"time"
)

// 并发世界中流行的模式之一是所谓的 fan-in 模式。这与 fan-out 模式相反，稍后我们将介绍。简而言之，fan-in 是一项功能，
//
//
//
// todo 可以从多个输入中读取数据并将其全部多路复用到单个通道中。

func Test_5(t *testing.T) {

	// 用main函数粘连两个chan
	ch := make(chan int)
	//out := make(chan int)

	go producer(ch, 100*time.Millisecond)
	go producer(ch, 100*time.Millisecond)

	go reader(ch)
	// todo 多路复用
	//for i := range ch {
	//	out <- i
	//}
	time.Sleep(1000 * time.Second)

}

func producer(ch chan int, d time.Duration) {

	var i int
	for {
		// 每隔三秒增加数据并写入通道中
		ch <- i
		i++
		time.Sleep(d)
	}

}

// 处理
func reader(out chan int) {
	for x := range out {
		fmt.Println(x)
	}
}
