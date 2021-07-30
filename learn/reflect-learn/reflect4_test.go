package map_test

import (
	"fmt"
	"testing"
	"time"
)

/*

复合结构体对象

*/

// []*Rectangle 如果直接当成数组传过去不会有改变
// 但是如果是加了指针传过去就会有改变
//
// map 是直接按指针传递的
func TestDo4(t *testing.T) {

	var maps map[string]*Rectangle = make(map[string]*Rectangle)
	go dosth(maps)
	rectangle := Rectangle{Heigh: 10}
	rectangle1 := Rectangle{Heigh: 11}
	maps["1"] = &rectangle
	maps["2"] = &rectangle1

	fmt.Println(fmt.Sprintf("REAL:%d", len(maps)))
	time.Sleep(3 * time.Second)
}
func dosth(maps map[string]*Rectangle) {
	time.Sleep(2 * time.Second)
	fmt.Println(len(maps))
}
