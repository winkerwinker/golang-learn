package goroutine_test

import (
	"fmt"
	"testing"
	"time"

)

/*
协程的使用
*/

func print() {
	fmt.Println("hello")
	time.Sleep(2 * time.Second)
}

func TestPool(t *testing.T) {

}
