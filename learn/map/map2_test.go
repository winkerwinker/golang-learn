package map_test

import (
	"testing"
)

type MyMap struct {
	Data map[int]string
	ch   chan func()
}






func TestMap2(t *testing.T) {

}


func NewMyMap() *MyMap {
	m := &MyMap{
		Data: make(map[int]string),
		// todo 这里的chan func 指的是什么
		// make(chan int) 指的是传入chan 指向的是int ，而这里的chan 指的向的是一个方法 ，一个方法会逐个吐出来执行
		ch:   make(chan func()),
	}

	go func() {
		for {
			// Redundant parentheses 冗余括号
			// todo 执行
			(<-m.ch)()
		}
	}()
	return m
}

func (m *MyMap) add(num int, data string) {
	m.ch <- func() {
		m.Data[num] = data
	}
}

func (m *MyMap) del(num int) {
	m.ch <- func() {
		delete(m.Data, num)
	}
}

func (m *MyMap) find(num int) (data string) {
	m.ch <- func() {
		if res, ok := m.Data[num]; ok {
			data = res
		}
	}
	return
}




