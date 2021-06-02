package map_test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

type Demo struct {
	Data map[string]string
	// 通用的锁
	Lock sync.Mutex
}

func (d Demo) Get(k string) string {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	return d.Data[k]
}

func (d Demo) Set(k, v string) {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	d.Data[k] = v
}

//fatal error: concurrent map read and map write
//还是出现了问题
func TestMap1(t *testing.T) {
	c := make(map[string]string)
	demo := Demo{Data: c, Lock: sync.Mutex{}}
	go func() {
		for j := 0; j < 1000000; j++ {

			demo.Set(strconv.Itoa(j), strconv.Itoa(j))
		}
	}()
	go func() {
		for j := 0; j < 1000000; j++ {
			fmt.Println(demo.Get(strconv.Itoa(j)))
		}
	}()

	time.Sleep(30 + time.Second)

}
