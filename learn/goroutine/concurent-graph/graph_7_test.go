package concurent_graph_test

import (
	"net"
	"testing"
)

// 下一个常见的模式类似于扇出，但是会在很短的时间内生成 goroutine，只是为了完成某些任务。它通常用于实现服务器 - 创建侦听器，循环运行 accept () 并为每个接受的连接启动 goroutine。它非常具有表现力，可以实现尽可能简单的服务器处理程序。看一个简单的例子：

// todo  每来一个创建一个协程来处理
func Test_8(t *testing.T) {

	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	for {

		accept, err := listen.Accept()
		if err != nil {
			continue
		}
		go handler(accept)
	}
}

func handler(c net.Conn) {
	c.Write([]byte("ok"))
	c.Close()
}
