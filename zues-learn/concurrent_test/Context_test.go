package context

import (
	"context"
	"fmt"
	"testing"
)

// context
//从设计的目的上说，context 是用来控制 goroutine 的，
//确切地说，是为了当你需要的时候，能干掉一个 goroutine。
//比如你用 go someFunc() 创建了成千上万个 goroutine，一切都很美好。突然你不想让它们继续执行了，
//这时你发现，糟了，你根本不知道这些 goroutine 放在哪，更别说控制它们了。
//不像其它语言用多线程时往往有个 Thread object，goroutine 并没有保存在某个变量里。怎么办？

func TestContext1(t *testing.T) {

	fields := []string{"1", "22", "33"}
	for _, field := range fields {
		field := field
		fmt.Println(field)
		go func() {
			fmt.Println(field)
		}()
	}

	//var uid string
	//fmt.Scanln(&uid)
	//fmt.Println("------ >")
	//fmt.Println(uid)
	//ctx, cancel := context.WithCancel(context.Background())
	//
	//go someFunc(ctx)
	//
	//time.Sleep(5 * time.Second)
	//
	//cancel()
}

func someFunc(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("endendend")
	}
}
