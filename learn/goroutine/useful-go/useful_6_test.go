package useful_go

import (
	"errors"
	"testing"
)

// 返回回调函数
func Test_6(t *testing.T) {

}

func bar() (string, error) {
	v, err := foo1()

	//直接进行返回
	if err != nil {
		return "", err
	}

	return v, nil

}

func bar1() (string, error) {
	return foo1()
}

func foo1() (string, error) {

	return "", errors.New("ok")
}
