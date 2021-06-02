package ttl

import (
	"fmt"
	"runtime"
	"time"
)

type expiringValue struct {
	expiration time.Time
	data       []byte //实际的值
}

type Map struct {
	data       map[string]expiringValue
	expiration time.Duration
}

func NewMap(expiration time.Duration) *Map {
	m := &Map{
		data:       make(map[string]expiringValue),
		expiration: expiration,
	}

	// start a worker goroutine
	go func() {
		for range time.Tick(expiration) {
			// todo map被回收了，但是工作协程还在运行并且还拥有一个对这个 map 实例的引用。
			// m.removeExpired()
		}
	}()

	return m
}

func main() {
	go func() {
		var stats runtime.MemStats
		for {
			runtime.ReadMemStats(&stats)
			fmt.Printf("HeapAlloc    = %d\n", stats.HeapAlloc)
			fmt.Printf("NumGoroutine = %d\n", runtime.NumGoroutine())
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		work()
	}
}

func work() {
	//m := NewMap(5 * time.Minute)
	//m.Set("my-key", []byte("my-value"))
	//
	//if _, ok := m.Get("my-key"); !ok {
	//	panic("no value present")
	//}
	// m超出变量范围

}
