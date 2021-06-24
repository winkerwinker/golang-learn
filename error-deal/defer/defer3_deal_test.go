package panics

import (
	"fmt"
	"testing"
)

// return 和 defer 之间的关系

//1. 返回变量 = xxx
//2. 调用 defer 函数（有可能更新返回变量的值）
//3. return 返回变量

//todo
//而如果没有定义，那么返回变量会是一个匿名的变量。defer 函数能够更改函数返回值的情况，都是在函数签名中定义了返回变量的情景。
func Test_4(t *testing.T) {

	//panic: runtime error: index out of range [0] with length 0 [recovered]
	//panic: runtime error: index out of range [0] with length 0
	// capcity 只能用于拷贝
	keys := make([]int64, 0, 100)
	keys[0] = 11
	fmt.Println(keys)
	//r := f1()
	//fmt.Println(r)
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}
