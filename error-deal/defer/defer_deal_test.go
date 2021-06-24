package panics

import (
	"fmt"
	"sync"
	"testing"
)

func Test_1(t *testing.T) {

	defer fmt.Println("defer")
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		go func() {
			wg.Add(1)
			fmt.Println("add 1")
			//time.Sleep(2 * time.Second)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("pre defer")
	//time.Sleep(1 * time.Minute)
}
