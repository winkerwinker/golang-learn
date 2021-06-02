package singeton_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var once sync.Once

// 单例模式
func TestOnces(t *testing.T) {

	for i, v := range make([]string, 10) {

		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}

	for i := 0; i < 10; i++ {

		go func() {
			once.Do(onced)
			fmt.Println(213)
		}()
	}

	time.Sleep(100)
}

func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
