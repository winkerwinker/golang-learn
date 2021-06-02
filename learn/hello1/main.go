package hello1

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

func hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println("hello world")
}
