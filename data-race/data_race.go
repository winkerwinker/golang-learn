package main

import (
	"fmt"
	"sync"
)

// 答案就是 data race tag，go 官方早在 1.1 版本就引入了数据竞争的检测工具，我们只需要在执行测试或者是编译的时候加上 -race 的 flag 就可以开启数据竞争的检测
//go test -race ./...
//go build -race

var wg sync.WaitGroup
var counter int

func main() {

	for i := 0; i < 100; i++ {
		run()
	}

}

func run() {
	for i := 0; i < 2; i++ {
		wg.Add(1)
		// 启动了两个协程
		go routine(i)
	}
	wg.Wait()
	fmt.Printf("final count %d \n", counter)
}

func routine(id int) {
	// 加上两次
	for i := 0; i < 2; i++ {
		value := counter
		value++
		counter = value
	}
	wg.Done()
}

/*

==================
WARNING: DATA RACE
Read at 0x00000064a9c0 by goroutine 7:
main.routine()
/main.go:29 +0x47

Previous write at 0x00000064a9c0 by goroutine 8:
main.routine()
/main.go:31 +0x64

Goroutine 7 (running) created at:
main.run()
/main.go:21 +0x75
main.main()
/main.go:14 +0x38

Goroutine 8 (finished) created at:
main.run()
/main.go:21 +0x75
main.main()
/main.go:14 +0x38
==================
exit status 66

*/



//这个结果非常清晰的告诉了我们在 29 行这个地方我们有一个 goroutine 在读取数据，但是呢，在 31 行这个地方又有一个 goroutine 在写入，所以产生了数据竞争。
//然后下面分别说明这两个 goroutine 是什么时候创建的，已经当前是否在运行当中。
// todo 前面带参数
// GORACE="halt_on_error=1 strip_path_prefix=/home/ll/project/Go-000/Week03/blog/03_sync/01_data_race" go run -race ./data_race.go


