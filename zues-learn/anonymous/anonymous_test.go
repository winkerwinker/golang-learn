package anonymous

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

//func()
//*reflect.rtype
func TestAnonymouswUse(t *testing.T) {
	//function variable
	fn := func() {
		fmt.Println("hello")
	}
	fmt.Printf("%T\n", fn)                 //打印func()
	fmt.Printf("%T\n", reflect.TypeOf(fn)) //打印func()
}

//func(int, int, float64) int
//*reflect.rtype
func TestAnonymouswUse1(t *testing.T) {
	fn := func(a, b int, z float64) int {
		return a * b
	}

	fmt.Printf("%T\n", fn) //打印fn的类型,
	fmt.Printf("%T\n", reflect.TypeOf(fn))
}

/*
defer func 可以推迟
go func 可以异步
*/
func TestAnonymouswUse2(t *testing.T) {

	// 也不是异步的
	//defer func(a, b int) {
	go func(a, b int) {
		time.Sleep(3 * time.Second)
		fmt.Println(a + b)
	}(1, 2)

	fmt.Println("------end")
	time.Sleep(4 * time.Second)
	//打印3
	// 后面作为参数传入

}
