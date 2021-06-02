package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("--------------")
	of := reflect.TypeOf("placeholder")

	// 为什么使用debug 查看会报错
	fmt.Println(of)

}


