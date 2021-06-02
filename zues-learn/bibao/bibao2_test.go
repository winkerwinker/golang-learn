package bibao

import (
	"fmt"
	"testing"
)

func TestContext3(t *testing.T) {
	// Q4实验：
	f7s := foo7(11)
	for _, f7 := range f7s {
		f7()

	}
}

func foo7(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		v := val
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+v)
		})
	}
	return fs
}
