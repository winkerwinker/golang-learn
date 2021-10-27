package context_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// 如果使用4个chanel
func TestContext5(t *testing.T) {

	oneWeekResultMap := make(map[int64]bool, 10)

	for i := 0; i < 10; i++ {
		oneWeekResultMap[1] = true

	}

	fmt.Println("------------------")
	bools := make(chan bool)

	// 只有一个能被关闭
	workChanel(bools, "node01")
	workChanel(bools, "node02")
	workChanel(bools, "node03")
	workChanel(bools, "node04")

	time.Sleep(5 * time.Second)
	fmt.Println("to stop")
	// 改动点
	bools <- false
	time.Sleep(10 * time.Second)

}

// 多个goroutine 被取消
func TestContext4(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	worker(ctx, "node01")
	worker(ctx, "node02")
	worker(ctx, "node03")
	worker(ctx, "node04")

	time.Sleep(5 * time.Second)
	fmt.Println("to stop")
	// 改动点
	cancel()
	time.Sleep(10 * time.Second)

}

// 1. 有多个goroutine 需要终止
// 2. goroutine 里面还有goroutine 需要终止
func TestContext3(t *testing.T) {

	// 多个goroutine 被取消
	// 最上层的节点
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			// 回传一个chanel， 如果调用了哪个canel ，这边done就会有值
			case <-ctx.Done():
				fmt.Println("stop")
				return
			default:
				fmt.Println("still")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("to stop")
	// 改动点
	cancel()
	time.Sleep(10 * time.Second)

}

// 使用 channel + select 主动停止
func TestContext2(t *testing.T) {
	stop := make(chan bool)

	go func() {
		for {
			// 控制器是一个通道 ，
			// 内部需要一个select 创建一个与控制器的链接
			select {
			case <-stop:
				fmt.Println("stop...")
				return
			default:
				fmt.Println("still")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	//
	time.Sleep(5 * time.Second)
	fmt.Println("to stop")
	stop <- true
	time.Sleep(10 * time.Second)
}

// 拆成很多个步骤跑 ，但是需要一起等待结果
// completeFuture Task
// 使用sync 包里面的waiting group
func TestContext1(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {

		time.Sleep(2 * time.Second)
		fmt.Println("j1 done")
		wg.Done()
	}()

	go func() {

		time.Sleep(1 * time.Second)
		fmt.Println("j2 done")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("all done")

}
