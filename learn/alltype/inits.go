package alltype

import (
	"fmt"
	"os"
)

/*
初始化类型群像
*/

func Init() {

	//const strings
	s := "hello"
	fmt.Println(s)

	if s[1] != 'e' {
		os.Exit(1)
	}

	s = "good bye"
	var p *string = &s

	*p = "ciao"

	// 拒绝再次赋值： Cannot assign to s[1]
	// s[1] = 'x'
	// (*p)[1] = 'y'

	// 因为改变了指向所以改变了数据源头
	fmt.Println(s)
}

func Init1() {
	// 申明数组 ，容量大了可以，容量小了会报错
	var arrayOfInt [10]int
	arrayOfInt = [10]int{1, 1, 1}
	// todo Println 会自动解析很多东西
	fmt.Println(arrayOfInt)

	sum(arrayOfInt[:])

}

func Init2() {
	// 四种初始化的方式
	ints := [10]int{1, 1, 2}
	ints1 := [...]int{1, 1, 2}
	slice := []int{1, 1, 2} //切片

	m := map[string]int{
		"s": 1,
		"a": 1,
		"b": 1,
	}

	_ = ints
	_ = ints1
	_ = slice
	_ = m

	// todo 理解这个结构
	for i := range ints {
		fmt.Println(i)
	}

}

// 运行for 循环，range关键字
// i, v 分别表示 下标和value
func Init3() {
	var list1 = []int{1, 3, 5, 7}
	for i := range list1 {
		fmt.Println(i)
		// 如果是i 对应的也是index
	}

	fmt.Println("------------")

	var list = [4]int{1, 3, 5, 7}
	for i, v := range list {
		fmt.Println(i, v)
	}

	// todo  go 中所有对应关系都用'：'表示，所有分割都用'，'表示
	m := map[string]int{
		"1": 1,
		"4": 1,
		"5": 1,
	}
	fmt.Println("------------")

	// 带上kv
	for s := range m {
		fmt.Println(s)
	}
	for s, k := range m {
		fmt.Println(s, k)
	}

	fmt.Println("------------")
	// 可以使用下划线，过滤掉没有用的值

	for _, s := range m {
		fmt.Println(s)
	}


	fmt.Println("------------")
	for i := 0; i < len(list1); i++ {
		fmt.Println(list1[i])
	}
}

// slice 作为参数，如果传入数组那么使用[:] 会自动转换
func sum(a []int) int { // returns an int
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
