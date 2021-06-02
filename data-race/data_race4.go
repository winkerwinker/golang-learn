package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)


//案例五 未受保护的成员变量
type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {

	atomic.StoreInt64(&w.last, time.Now().UnixNano())
	// w.last = time.Now().UnixNano() // First conflicting access.
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// Second conflicting access.
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
			// if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}

//
//这是因为我们在 maker = jerry
//这种赋值操作的时候并不是原子的，
//在上一篇文章中我们讲到过，
//只有对 single machine word
//进行赋值的时候才是原子的，虽然这个看上去只有一行，但是 interface 在 go 中其实是一个结构体，它包含了 type 和 data 两个部分，所以它的复制也不是原子的，会出现问题
// todo  只有对single machine word 赋值才是院子的



// dealntodo https://lailin.xyz/post/go-training-week3-data-race.html