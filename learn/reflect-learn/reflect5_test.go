package map_test

import (
	"fmt"
	"testing"
)

/*

复合结构体对象

*/

func TestDo5(t *testing.T) {
	fmt.Printf("%b\n", 1<<4)
	fmt.Printf("%b\n", 1<<5)
	fmt.Println(1 << 4)
	fmt.Printf("%b\n", 1<<5&1<<4)

	fmt.Println(1 << 5 & 1 << 5)
	fmt.Println(1 << 5 & 1 << 4)
	fmt.Println(1 << 4 & 1 << 5)

	fmt.Println(flagEmbedRO)
	fmt.Println(1 << 2)
	fmt.Println(1 << 3)
	fmt.Println(1 << 4)

	fmt.Println(1 << 6)
	fmt.Println(1<<0 | 1<<1)
	fmt.Println(1 << 0 & 1 << 1)

	fmt.Println(1 << 0) // 01
	fmt.Println(1 << 1) // 10

}

type flag uintptr

const (
	flagKindWidth        = 5 // there are 27 kinds
	flagKindMask    flag = 1<<flagKindWidth - 1
	flagStickyRO    flag = 1 << 5
	flagEmbedRO     flag = 1 << 6
	flagIndir       flag = 1 << 7
	flagAddr        flag = 1 << 8
	flagMethod      flag = 1 << 9
	flagMethodShift      = 10
	flagRO          flag = flagStickyRO | flagEmbedRO
)
