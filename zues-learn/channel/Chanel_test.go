package chanel

import (
	"fmt"
	"testing"
)

//往一个已经被close的channel中继续发送数据会导致run-time panic。
//往nil channel中发送数据会一致被阻塞着。
func TestChanel1(t *testing.T) {

	c := make(chan int)
	defer close(c)
	// 写入chanel 在写入i
	go func() { c <- 3 + 4 }()
	x, ok := <-c
	fmt.Println(ok)
	fmt.Println(x)

	//fatal error: all goroutines are asleep - deadlock!,
	//todo 只是再make了一个chanel而已
	nilChanel := make(chan int)
	x, ok = <-nilChanel
	fmt.Println(ok)
	fmt.Println(x)
	//x, ok := <-c
	//all goroutines are asleep - deadlock!

}
