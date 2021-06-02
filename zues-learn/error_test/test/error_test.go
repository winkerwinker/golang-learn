package test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

// 测试五种错误处理的方式

func TestError(t *testing.T) {



}

//传播错误: 将错误返回给上级，变成这个函数整体的失败
func fun() (resp *http.Response, err error) {

	get, err := http.Get("htt")
	if err != nil {
		return nil, err
	}
	return get, err
}

// retry
func waitForserver(url string) error {
	const minute = time.Minute
	ddl := time.Now().Add(minute)

	// +1 直到超时
	for tries := 1; time.Now().Before(ddl); tries += 1 {
		_, err := http.Get("htt")
		if err == nil {
			return nil
		}

		// 错误并且睡眠
		log.Printf("Server not respond (%v);the %d times retry....", err, tries)
		time.Sleep(time.Second * 2)

	}
	return fmt.Errorf("connot connect to server")
}

// 输入信息并结束  ->  如何应用到一个rpc上，rpc可以提前结束吗
func waitForserver1(url string) {
	os.Exit(1)
}


// 忽略错误