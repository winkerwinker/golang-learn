package nil

import "fmt"

/*
申明变量使用var a string
*/

func Demo(name string) {

	var pi *int = nil
	var i interface{}
	// todo 为什么可以把指向int 的指针 指向 interface{}
	i = pi
	fmt.Print(i == nil)

}

/*
未使用的报错 ，可以使用虚拟使用解决
*/
func unuse() {
	var n int = 5
	// todo
	_ = n
}

//todo 为什么java不带有指针会更加好扫描root 一些 （所谓的拼接是什么）