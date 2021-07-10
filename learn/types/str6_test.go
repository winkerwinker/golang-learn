package cache_sdk

import (
	"fmt"
	"testing"
)

var (
	lc3Errcnt = 0
)

func Test_test(t *testing.T) {
	Lc3Chan := make(chan int)

	// 就阻塞在这里等待了，所以所有的阻塞都要用go
	go func() {
		collectAllOzUpdateAndErrCnt(Lc3Chan)
	}()

	fmt.Println(&Lc3Chan)

	for i := 0; i < 11; i++ {
		Lc3Chan <- i
	}

}

func collectAllOzUpdateAndErrCnt(chans chan int) {
	fmt.Println(&chans)

	for i := 0; i < 11; i++ {

		fmt.Println(lc3Errcnt)

		// 一个消费一个
		// 接下去就会被阻塞在这里
		select {
		// todo 如果有两个同样chanel会随机进一个
		case cnt := <-chans:
			fmt.Println("cnt1")
			lc3Errcnt += cnt
		case cnt := <-chans:
			fmt.Println("cnt2")
			lc3Errcnt += cnt
		}
		// 两个chanel 都会被进入
		fmt.Println(lc3Errcnt)

	}
}
