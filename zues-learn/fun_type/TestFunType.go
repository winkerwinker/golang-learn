package fun_type

import (
	"fmt"
	"math"
	"testing"
)

func TestFun(t *testing.T) {
	fmt.Println("嗨客网(www.haicoder.net)")


}


//func getPageCount(pageSize int) float64 {
//	if pageSize <= 0{
//		return
//	}
//
//	pageCount := math.Ceil(float64(100/pageSize))
//	return pageCount
//}


//(int float64) {  如果这里加了就会有默认值
func getPageCount1(pageSize int) (int float64) {
	if pageSize <= 0{
		return
	}

	pageCount := math.Ceil(float64(100/pageSize))
	return pageCount
}
