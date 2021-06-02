package main

import (
	"awesomeProject/learn/handler"
	"fmt"
)

func main() {
	//// handler.PeInfo()
	//people := handler.People{Name: "xieaichen", CurrStep: 0}
	//people.PeInfo()
	//// 不会改变原有的数据
	//fmt.Println(people.CurrStep)
	//
	//people1 := handler.People{Name: "xieaichen", CurrStep: 0}
	//people1.PeInfo1()
	//// 会改变原有的数据
	//fmt.Println(people1.CurrStep)


	peopl2 := handler.People{Name: "xieaichen", CurrStep: 0}
	slot := &peopl2
	slot.PeInfo()
	// 不会改变原有的数据
	fmt.Println(peopl2.CurrStep)
}
