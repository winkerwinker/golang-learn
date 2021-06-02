package json

import (
	"fmt"
	"testing"
)

func TestSay2(t *testing.T) {

	var f float64 = 1.23
	s := demo(f)
	fmt.Println(s)
}

func demo(inter interface{}) float64 {

	return inter.(float64)

}
