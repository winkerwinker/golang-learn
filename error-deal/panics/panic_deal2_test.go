package panics

import (
	"fmt"
	"testing"
)

func Test_3(t *testing.T) {

	value1 := make([]interface{}, 1)
	value1[0] = nil
	fmt.Println(value1)
	i2 := interface{}(value1)
	fmt.Println(i2)
	i := i2.([]interface{})[0]
	fmt.Println(i)
}
