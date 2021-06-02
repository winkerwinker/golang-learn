package json

import (
	"fmt"
	"testing"
)

func TestSay4(t *testing.T) {

	result := make([]float64, 3, 3)
	result[0]=0
	result[1]=1
	result[2]=2
	result[3]=3

	fmt.Println(result)
	fmt.Println("------")
	result1 := []float64{1000.0, 2.0, 3.4, 7.0, 50.0}
	result2 := []float64{0, 1, 2, 3, 4}
	result3 := []float64{9, 8, 7, 6, 5}
	reuslt11 := [][]float64{result1, result2}
	reuslt12 := [][]float64{result1, result3}
	reuslt13 := [][]float64{result2, result3}
	reuslt111 := [][][]float64{reuslt11, reuslt12, reuslt13}
	fmt.Println(reuslt111)
	scape := mergeLandScape(reuslt111)
	fmt.Println(scape)

	fmt.Println("---------")
	scape1 := merge(reuslt11, reuslt12, reuslt13)
	fmt.Println(scape1)
}

func mergeLandScape(list [][][]float64) (result [][]float64) {
	for index, _ := range list[0] {
		float64s := make([]float64, 0, 10)
		for in, _ := range list {
			float64s = append(float64s, list[in][index]...)
		}
		result = append(result, float64s)
	}
	return
}

func merge(result1, result2, result5 [][]float64) [][]float64 {
	for in, date := range result2 {
		result1[in] = append(result1[in], date...)
	}
	for in, date := range result5 {
		result1[in] = append(result1[in], date...)
	}

	return result1
}
