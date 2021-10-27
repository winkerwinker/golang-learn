package main

import "fmt"

// go build -gcflags="-m -l" main1.go

func main() {
	// 每一个变量的创建都会有通报是否会内存逃逸

	s := "ok"
	test1(s) //凡是传interface 在这个时候逃逸了

	//因为，Golang中，slice,map,channel引用指针的变量，一定会逃逸。 Golang中，slice，map，channel对指针的引用会比之保留变量的slice，map，channel性能低，这里是根本原因。
	//todo  golang最好直接保存变量 ，因为指针会逃逸
	a := make([]*int, 1)
	b := 12
	a[0] = &b

	c := make(map[string]*int)
	d := 14
	c["aaa"] = &d

	e := make(chan *int, 1)
	f := 15
	e <- &f
}

func test1(str interface{}) {

	fmt.Println(str)
}
