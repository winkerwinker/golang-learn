package useful_go

import (
	"fmt"
	"testing"
	"time"
)

//将 for-select 语法结构封装成函数#
func Test_1(t *testing.T) {

	// blocking
	if err := foo(); err != nil {
		// 处理错误
	}

	fmt.Println("ending")

}

func foo() error {
	for {
		select {
		case <-time.After(time.Second):
			fmt.Printf("hello")

		default:
			break
		}
	}

}

// todo 将 for 包装成一个函数来处理业务逻辑
//	for {
//		select {
//		case <-time.After(time.Second):
//			fmt.Printf("hello")
//
//		default:
//			break
//		}
//	}
