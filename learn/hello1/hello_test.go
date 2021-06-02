package hello1

import (
	"fmt"
	"testing"
)

/*
测试需要满足三个要点

程序需要在一个名为 xxx_test.go 的文件中编写
测试函数的命名必须以单词 Test 开始
测试函数只接受一个参数 t *testing.T

类型为 *testing.T 的变量 t 是你在测试框架中的 hook（钩子）

*/


func TestHello(t *testing.T) {
	// 一定会引用相对应的变量吗
	got := hello("acxie")
	want := englishHelloPrefix

	fmt.Println(got)
	fmt.Println(want)

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}

	fmt.Println("tests")
}


func TestHelloWorld(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		// 申明是辅助函数
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := hello("Chris")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}