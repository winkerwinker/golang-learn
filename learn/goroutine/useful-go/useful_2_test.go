package useful_go

import (
	"fmt"
	"testing"
)

type T struct {
	Foo string
	Bar int
}

// 在初始化结构体的时候，使用有标签的语法
func Test_2(t *testing.T) {

	// todo 如果结构体变化，编译就会失败
	t2 := T{"example", 123} // 无标签语法
	//t %+vn 是啥
	fmt.Printf("t %+vn", t2)

}
