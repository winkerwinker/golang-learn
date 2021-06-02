package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	//在新启动的 goroutine 当中读取 i 的值，在 main 中写入，导致出现了 data race
	for i := 0; i < 5; i++ {
	go func() {
		fmt.Println(i) // Not the 'i' you are looking for.
		wg.Done()
	}()
}
wg.Wait()
}
