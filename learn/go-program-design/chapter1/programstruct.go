package main

import (
	"errors"
	"fmt"
)

/*
程序规范:
包名字使用小写
变量命名为驼峰
文件名呢？
*/
/*
变量的声明：var name type = expression
*/

func initvar() {
	//var i, j, k int
	//var b, f, s = true, 2.3, "four"
	//:= 短变量声明
}

func shortVar() {
	var err error
	// 在这个局域内就是初始化变量
	if f, err := result(); err != nil {
		fmt.Println(f)
	}
	fmt.Println(err)
}

func result() (string, error) {
	return "", errors.New("text")
}
