package main

import (
	"fmt"
	"testing"
)

type man struct {
	Name string
	Sex  string
}

type geometry interface {
	// 实现了接口的方法就是实现了接口
	say() float64
}

// todo
func (m *man) say() {
	fmt.Println("-----")
	sayOut(m.Name)
	fmt.Println("-----")
}

func sayOut(a string) {
	fmt.Println(a)
}

func TestSay(t *testing.T) {
	// 可以这样直接传递指针


	// m := man{Name: "hello", Sex: "nv"}
	// 不能给对象直接设置为空指针
	// var m man = nil


	// m := &man{Name: "hello", Sex: "nv"}
	 var m *man = nil
	// 如果这是 m 是nil ，可以继续走调用吗
	m.say()
}