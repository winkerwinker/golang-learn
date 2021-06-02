package main

import (
	"fmt"
)

/*

todo  命名规则
1. 文件  -> _分割
1. 方法  -> 大驼峰
1. 变量  -> 小驼峰


支持两个参数

1. 指定问候的人。
1. 指定问候的语言。如果一种不能识别的语言被传进来，就默认为英语。

*/
const englishHelloPrefix = "Hello, "
const defaultLanguage = "English "

func Hello(name string, language string) string {

	if language == "" {
		language = defaultLanguage
	}
	return englishHelloPrefix + name + "使用" + language
}

func main() {
	fmt.Println("hello world")
}
