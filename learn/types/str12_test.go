package cache_sdk

import (
	"fmt"
	"testing"
)

func Test_create12(t *testing.T) {
	_, err := getResult()
	if err == nil {
		fmt.Printf("oooo")
		return
	}
	defer deferFunc(err)
}

func getResult() (string, error) {
	return "", nil
}
func deferFunc(err error) {
	fmt.Printf("oooo")
	fmt.Println(err)
	fmt.Printf("oooo")
}
