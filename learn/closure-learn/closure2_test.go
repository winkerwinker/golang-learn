package map_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

/*
创建复杂类型的对象
可以像原始对象一样使用零值来创建复合对象。但是，仅使用一个空值而不做其他额外的操作的话，它们并不能被初始化。下面一节将详细讨论如何初始化复合种类。

*/

var confListener = make(chan int)

func TestDo2(t *testing.T) {

	go listen(context.Background())
	time.Sleep(10 * time.Second)
	confListener <- 1
}

func TestDo21(t *testing.T) {

	 fmt.Println(2 << 4)

	//NewRedisClient("Test", []string{"10.190.13.75:9000"}, 10, 10, time.Second*10, time.Second*10, time.Second*10, time.Second*10)

	source := ApolloDataSource{NameSpace: "CXYX-Zeus", Name: "basic_20201109",Type: "apollo"}
	marshal, _ := json.Marshal(source)
	fmt.Println(string(marshal))

}

// apollo 数据源，没有其他任何附加的映射方式
type ApolloDataSource struct {
	NameSpace        string
	Name             string
	Type             string
	nodes            []string
	maxActive        int
	maxIdle          int
	idleTimeout      time.Duration
	dialConnTimeout  time.Duration
	dialReadTimeout  time.Duration
	dialWriteTimeout time.Duration
}

// todo 注意这种写法
func listen(ctx context.Context) {
	for {
		select {
		case <-confListener:
			{
				fmt.Println("ok")
			}
		case <-ctx.Done():
			return
		}

		fmt.Println("daole")
	}
}
