package gonum_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)
import "gonum.org/v1/gonum/mat"

// todo 处理NaN
func calculateScore(mtl []float64, pctrcvrIndex int, saleCntIndex int, min float64, max float64, price float64, a float64, b float64, c float64, d float64) float64 {
	return math.Pow(mtl[pctrcvrIndex], a) + math.Pow((math.Log(Clip(mtl[saleCntIndex], min, max)*price*c+d)), b)

	clip := Clip(mtl[saleCntIndex], min, max)
	fmt.Printf("clip : %v\n", clip)

	f := clip*price*c + d
	fmt.Printf("f : %v\n", f)

	log := math.Log(f)
	fmt.Printf("log : %v\n", log)

	pow := math.Pow(mtl[pctrcvrIndex], a)
	fmt.Printf("pow : %v\n", pow)

	pow2 := math.Pow(log, b)
	fmt.Printf("pow2 : %v\n", pow2)
	return pow + pow2
}

func Clip(x float64, min float64, max float64) float64 {

	if x < min {
		return min
	}

	if x > max {
		return max
	}

	return x
}

func Test(t *testing.T) {

	//	a:1
	//b:0.2
	//c:1
	//d:1
	//min:0.001
	//max:3

	a, _ := strconv.ParseFloat("1", 64)
	b, _ := strconv.ParseFloat("0.2", 64)
	c, _ := strconv.ParseFloat("1", 64)
	d, _ := strconv.ParseFloat("1", 64)
	min, _ := strconv.ParseFloat("0.001", 64)
	max, _ := strconv.ParseFloat("3", 64)

	//time="2021-10-28T15:39:14+08:00" level=info msg=_com_debug  371009230  ctrcvr=0.006088367663323879 price=3.1 saleCnt=0.00934762042015791 score=NaN
	//time="2021-10-28T15:39:14+08:00" level=info msg=_com_debug 371011255 ctrcvr=0.005832429975271225 price=2 saleCnt=0.008596064522862434 score=NaN

	score := calculateScore([]float64{0.006088367663323879, 0.00934762042015791}, 0, 1, min, max, 3.1, a, b, c, d)
	fmt.Printf("score : %v\n", score)

	//funcName()
	//	2, 6, 8, -4,
	//	1, 8, 7, -2,
	//	2, 2, 1, 7,
	//	8, -2, -2, 1,
	//})
	//
	//// 高维矩阵
	////mat.VecDense{}
	//
	//matrix := tmp.T()
	//fmt.Println(matrix)
	//
	//formatted := mat.Formatted(matrix, mat.FormatPython())
	//fmt.Println(formatted)

	//components := []float64{1.2, -5.7, -2.4, 7.3}
	//a := mat.NewDense(2, 2, components)
	//fa := mat.Formatted(a, mat.FormatPython())

}

func funcName() {
	i := [][]float64{
		[]float64{2, 6, 8, -4},
		[]float64{1, 8, 7, -2},
		{2, 2, 1, 7},
		{8, -2, -2, 1},
	}

	dense := getMatrix(i)
	formatted := mat.Formatted(dense, mat.FormatPython())
	fmt.Println("-------")
	fmt.Println(formatted)
	fmt.Println("-------")
	//tmp := mat.NewDense(4, 4, []float64{
}

//多维度数组转 matrix
func getMatrix(input [][]float64) *mat.Dense {
	tmpInput := make([]float64, 0, len(input)*len(input[0]))

	for i := 0; i < len(input); i++ {
		tmpInput = append(tmpInput, input[i]...)
	}
	return mat.NewDense(len(input), len(input[0]), tmpInput)

}
