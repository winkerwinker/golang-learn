package cache_sdk

import (
	"fmt"
	"testing"
	"time"
)

func Test_create10(t *testing.T) {
	testSend()
	time.Sleep(3 * time.Second)
	fmt.Println("done")
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic2 at :")
				fmt.Println(err)
			}
		}()
	}()
}

func testSend() {
	var (
		errChan = make(chan string, 10)
	)
	// panic 处理
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic at :")
			}
		}()
	}()

	defer close(errChan)

	// panic("sorry!!")

	// 主要业务
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("hello")
			if i == 3 {
				panic("sorry!!")
			}
			time.Sleep(20 * time.Millisecond)
			errChan <- "hello"
		}(i)
	}

	for i := 0; i < 10; i++ {
		select {
		case <-time.After(time.Millisecond * 75):
		case errorStr := <-errChan:
			{
				fmt.Println(errorStr)
			}
		}
	}

	panic("sorry!!")

}
