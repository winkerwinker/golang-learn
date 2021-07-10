package cache_sdk

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_bugfix(t *testing.T) {
	//array := ConvertFloat64StringToArray("[889]")
	//fmt.Println(array)
	m := map[string]string{"11": "11"}
	s := m["22"]
	fmt.Println("====")
	fmt.Println(s)
	if s == "" {
		s = "-1"
	}
	fmt.Println("====")
	fmt.Println(s)
	//array = ConvertFloat64StringToArray("[]")
	//fmt.Println(array)
}

func ConvertFloat64StringToArray(s string) []float64 {
	a := strings.Split(s, ",")
	fmt.Println("=========")
	fmt.Println(a)
	fmt.Println("=========")
	if len(a) == 0 {
		return []float64{}
	}
	res := make([]float64, len(a))
	for index, value := range a {

		if index == 0 {
			value = strings.Replace(value, "[", "", 1)
		}
		if index == len(a)-1 {
			value = strings.Replace(value, "]", "", 1)
		}
		if floatvalue, err := strconv.ParseFloat(value, 64); err != nil {
			res[index] = 0.0
		} else {
			res[index] = floatvalue
		}
	}
	return res
}
