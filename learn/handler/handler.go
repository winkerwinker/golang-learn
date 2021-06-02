package handler

import "fmt"

// 接口
type Animal interface {
	Walk(step int)
	Walk1(step int)
}

// 类
type People struct {
	Name     string
	CurrStep int
}

func (c People) PeInfo() {
	c.CurrStep++
	fmt.Println("people name:", c.Name, " currstep:", c.CurrStep)
}

func (c *People) PeInfo1() {
	c.CurrStep++
	fmt.Println("people name:", c.Name, " currstep:", c.CurrStep)
}
