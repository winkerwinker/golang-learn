package map_test

import (
	"fmt"
	"testing"
)

//fatal error: PASS
//concurrent map read and map write
//多个协程同时写也会出现fatal error: concurrent map writes的错误
// map 读写控制

type Filetemp struct {
	fd   int    // file descriptor number
	name string // file name at Open time
}

func TestMap(t *testing.T) {

	c := make(map[string]*Filetemp)
	go func() { //开一个协程写map
		for j := 0; j < 1000000; j++ {
			c[fmt.Sprintf("%d", j)] = &Filetemp{fd: 1, name: "test"}
		}
	}()
	go func() { //开一个协程读map
		for j := 0; j < 1000000; j++ {
			filetemp := c[fmt.Sprintf("%d", j)]
			// 会出现并发问题
			filetemp.name="hhhh"
		}
	}()

	// time.Sleep(time.Second * 20)

	//test := make(map[string]int)
	//test["1"] = 1
	//test["88"] = 2
	//
	//
	//for i := 0; i < 10000; i++ {
	//	test[strconv.Itoa(i)]= 5
	//	go  fmt.Sprintf("%q", test)
	//}
	//
	//
	//time.Sleep(9 * time.Second)
	// fmt.Println(sprintf)
	//c := make(map[string]Filetemp)
	//
	//go func() {
	//	for j := 0; j < 10; j++ {
	//		c[fmt.Sprintf("%d", j)] = Filetemp{fd: 1, name: "test"}
	//	}
	//}()
	//
	//time.Sleep(2 * time.Second)
	//i := c["88888"]
	//fmt.Println(i)
	//go func() {
	//	for j := 0; j < 1000000; j++ {
	//		fmt.Println(c[fmt.Sprintf("%d",j)])
	//	}
	//}()
}
