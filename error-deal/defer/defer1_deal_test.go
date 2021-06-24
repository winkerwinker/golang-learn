package panics

import (
	"fmt"
	"testing"
)

func Test_2(t *testing.T) {
	a := increaseAAA()
	fmt.Println(a)
}

// 1. 参数传递
//作为 函数参数，则在 defer 定义时 就把值传递给 defer，并被 缓存 起来；
//作为 闭包引用 的话，则会在 defer 函数真正调用时根据整个上下文确定当前的值。
// 2. 执行顺序
// 当前函数返回前再把延迟函数取出并执行
//defer 定义的函数会先进入一个栈，函数 return 前，会按先进后出（FILO）的顺序执行。也就是说最先被定义的 defer 语句最后执行。
func increaseA() int {
	var i int
	//defer 语句定义时，对 外部变量的引用 是有两种方式的，
	//分别是作为 函数参数 和作为 闭包引用。
	defer func() {
		i++
	}()
	// 在这个位置会执行defer
	return i
}

// todo 虽然作为参数穿进去了，但是改变的是内部的引用，而不是指针
func increaseAA() int {
	var i int
	//defer 语句定义时，对 外部变量的引用 是有两种方式的，
	//分别是作为 函数参数 和作为 闭包引用。
	defer func(index int) {
		index++
	}(i)
	// 在这个位置会执行defer
	return i
}

// todo 直接修改指针为什么无效
func increaseAAA() int {
	var i int
	// 实参在 defer 函数定义时就取值并缓存起来了。
	defer func(index *int) {
		(*index)++
	}(&i)
	// 在这个位置会执行defer
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func increaseC() (r int) {
	defer func(r int) {
		r++
	}(r)
	return r
}
