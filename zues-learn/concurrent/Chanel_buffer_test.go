package chanel

import (
	"fmt"
	"testing"
	"time"
)

//Go 缓冲 channel 和 非缓冲 channel 的区别
func TestChanelBuffer(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1 //写进去了就是写进去了
	ch <- 2
	ch <- 3
	go loop(ch)
	time.Sleep(1 * time.Millisecond)

}

func loop(ch chan int) {
	for {
		i := <-ch
		fmt.Println("this  value of unbuffer channel", i)
	}

}

/*
阻塞：在执行过程中暂停，以等待某个条件的触发 ，我们就称之为阻塞
在Go中我们make一个channel有两种方式：分别是有缓冲的和没缓冲的
缓冲channel 即 buffer channel 创建方式为
1.
make(chan TYPE,SIZE)
如make(chan int,3) 就是创建一个int类型，缓冲大小为3的 channel
2.
非缓冲channel 即 unbuffer channel 创建方式为 make(chan TYPE)
如 make(chan int) 就是创建一个int类型的非缓冲channel

3.
非缓冲channel 和 缓冲channel 的区别
非缓冲 channel，channel 发送和接收动作是同时发生的
例如 ch := make(chan int) ，如果没 goroutine 读取接收者<-ch ，那么发送者ch<- 就会一直阻塞

*/
// todo 缓冲 channel 类似一个队列，只有队列满了才可能发送阻塞
