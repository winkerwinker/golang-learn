package chanel

import (
	"fmt"
	"testing"
)

// 用于blocking阻塞
// Buffered Channels 指的是缓存通道吗
// go中阻塞的特性，直接可以用来协调两个线程的顺序，但是java中需要用锁、信号量等一系列工具操作
func TestChanel2(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func demo(s []int, c chan int) {
	// 发送和接收的语法
	// 在通道上箭头的方向指定数据是发送还是接收。
	a := make(chan string)
	data := <-a // read from channel a
	a <- data   // write to channel a
}
