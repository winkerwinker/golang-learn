package main

import (
	"fmt"
	"io/ioutil"
)


func main() {
	// 平常处理
	fmt.Println("----------")
	ordinary()
}





type fileError struct {
}

func (fe *fileError) Error() string {
	return "文件错误"
}



func ordinary() {
	conent, err := ioutil.ReadFile("filepath")
	// 处理原则 ： 要么你处理、要么你忽略
	if err != nil {
		//错误处理
		fmt.Println(err)
	} else {
		fmt.Println(string(conent))
	}
}


