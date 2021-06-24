package panics

import (
	"fmt"
	"testing"
)

func Test_3(t *testing.T) {
	point()
}

func trace(str string) string {
	// 1. 最先执行
	fmt.Println("entering " + str)
	return str
}
func leave(str string) {
	// 3. 执行语句
	fmt.Println("leaving " + str)
}
func point() {
	defer leave(trace("point"))
	// 2. 执行其他语句
	fmt.Println("in point")
}
