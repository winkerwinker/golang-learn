package map_test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

/*
原子计数器
*/

func TestDo1(t *testing.T) {
	//fmt.Println("=======unsafe========")
	//countWithThreadUnsafe()
	ints := []int{1, 2, 3}
	i := ints[len(ints):]
	fmt.Println(i)
	//fmt.Println("=========safe========")
	//countWithThreadSafe()

}

func countWithThreadSafe() {
	var ops uint64 = 0
	//for i := 0; i < 1000; i++ {
	go func() {
		for {
			// To atomically increment the counter we
			// use `AddUint64`, giving it the memory
			// address of our `ops` counter with the
			// `&` syntax.
			atomic.AddUint64(&ops, 1)
			swapUint64 := atomic.CompareAndSwapUint64(&ops, 10, 0)

			fmt.Println("tmp ops:", atomic.LoadUint64(&ops), "swapped: ", swapUint64)
			// Wait a bit between increments.
		}
	}()
	//}

	time.Sleep(3 * time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

func countWithThreadUnsafe() {

	var ops uint64 = 0
	//for i := 0; i < 1; i++ {
	go func() {
		for {
			//ops = ops + 1
			// 普通的ops 不会刷新到主内存去，而AddUint64会强制花刷新到主内存
			atomic.AddUint64(&ops, 1)
			//time.Sleep( time.Millisecond)
			//fmt.Println("ops:", ops)
		}
	}()
	//}
	fmt.Println("final ops:", ops)
	time.Sleep(3 * time.Second)

	fmt.Println("final ops:", ops)
}
