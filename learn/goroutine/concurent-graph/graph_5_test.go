package concurent_graph_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//  todo 消息过来，我开多个协程给你处理

//与 fan-in 相反的模式是 fan-out 或者 worker 模式。
//多个 goroutine 可以从单个通道读取，
//从而在 CPU 内核之间分配大量的工作量，
//因此是 worker 的名称。在 Go 中，此模式易于实现 - 只需以通道为参数启动多个 goroutine，
//然后将值发送至该通道 - Go 运行时会自动地进行分配和复用

func Test_6(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(4)
	// fatal error: all goroutines are asleep - deadlock!

	// 当所有协程都关闭了之后， 但是wg还没有done
	go pool(&wg, 3, 5)
	// Wait blocks until the WaitGroup counter is zero.
	wg.Wait()

}

// 工作线程
func woker(taskCh <-chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for {
		ch, ok := <-taskCh
		if !ok {
			return
		}
		//  睡一会
		d := time.Duration(ch) * time.Millisecond
		time.Sleep(d)

		fmt.Println("processing task", d)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {

	taskCh := make(chan int)

	// 启动一定的工作线程
	for i := 0; i < workers; i++ {
		go woker(taskCh, wg)
	}

	// 将任务放进工作线程
	for i := 0; i < tasks; i++ {
		taskCh <- i
	}

	close(taskCh)

	fmt.Println("---------chan close")

}

func Test_60(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(2)
	// fatal error: all goroutines are asleep - deadlock!

	// 当所有协程都关闭了之后， 但是wg还没有done
	//for i := 0; i < 2; i++ {
	//	// 必须传入指针， 不然操作无用
	//	go func(wg *sync.WaitGroup) {
	//		fmt.Println("done once")
	//		wg.Done()
	//	}(&wg)
	//}

	for i := 0; i < 2; i++ {
		// 必须传入指针， 不然操作无用
		go func(wg sync.WaitGroup) {
			fmt.Println("done once")
			wg.Done()
		}(wg)
	}

	// time.Sleep(4 * time.Second)

	// Wait blocks until the WaitGroup counter is zero.
	wg.Wait()

}

func Test_600(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(2)
	// fatal error: all goroutines are asleep - deadlock!

	// 当所有协程都关闭了之后， 但是wg还没有done
	wg.Done()
	// 如果只有一个done会有问题
	wg.Done()

	// Wait blocks until the WaitGroup counter is zero.
	wg.Wait()

}
