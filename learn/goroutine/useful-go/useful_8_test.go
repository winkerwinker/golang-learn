package useful_go

import (
	"sync"
	"testing"
)

// withContext
func Test_8(t *testing.T) {

}

var mu sync.Mutex

// 将重复方法包函数， 自由调用的函数作为参数

func withLockContext(fn func()) {
	mu.Lock()
	defer mu.Unlock()

	fn()
}

//有时对于函数会有一些重复劳动，例如锁 / 解锁，初始化一个新的局部上下文，准备初始化变量等等…… 这里有一个例子：

// todo 封装完成
func foo3() {
	withLockContext(func() {
		// foo 相关的工作
	})
}

func foo2() {
	mu.Lock()
	defer mu.Unlock()

	// foo 相关的工作
}

func bar2() {
	mu.Lock()
	defer mu.Unlock()

	// bar 相关的工作
}

func qux2() {
	mu.Lock()
	defer mu.Unlock()

	// qux 相关的工作
}
