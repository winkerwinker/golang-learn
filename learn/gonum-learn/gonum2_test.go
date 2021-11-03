package gonum_test

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"reflect"
	"testing"
	"time"
)

//多维度数组转 matrix
func GetMatrix(input [][]float64) *mat.Dense {
	tmpInput := make([]float64, 0, len(input)*len(input[0]))

	for i := 0; i < len(input); i++ {
		tmpInput = append(tmpInput, input[i]...)
	}
	return mat.NewDense(len(input), len(input[0]), tmpInput)

}

//多维度数组转 matrix
func GetMatrixZero(input [][]float64) *mat.Dense {
	return mat.NewDense(len(input), len(input[0]), nil)

}

func TestLinUCBInference(t *testing.T) {

	resCh := make(chan interface{}, 8)
	//errCh := make(chan error, 1)

	fmt.Println(len(resCh))

	go func() {
		for i := 0; i < 5; i++ {
			resCh <- 1
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(resCh))

	type args struct {
		A        [][]float64
		b        []float64
		features []float64
	}
	tests := []struct {
		name    string
		args    args
		wantS   fmt.Formatter
		wantErr bool
	}{
		{name: "test1",
			args:    args{A: [][]float64{{1, 12, 8}, {0, 4, 5}, {0, 0, 9}}, b: []float64{0.1, 0.6, 0.25}, features: []float64{0, 1, 1}},
			wantS:   nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := LinUCBInference(tt.args.A, tt.args.b, tt.args.features)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinUCBInference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("LinUCBInference() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

// 包装乘法方法
func Multi(a, b mat.Matrix) *mat.Dense {
	tmpDense := mat.NewDense(GetMatrixRow(a), GetMatrixCol(b), nil)
	tmpDense.Mul(a, b)
	return tmpDense
}

func GetMatrixRow(b mat.Matrix) int {

	i, _ := b.Dims()
	return i

}

func GetMatrixCol(b mat.Matrix) int {
	_, j := b.Dims()
	return j
}

// 包装乘法方法
func Inverse(a mat.Matrix) (*mat.Dense, error) {
	i, j := a.Dims()
	tmp := mat.NewDense(i, j, nil)
	// AMatrixInverse := GetMatrixZero(A)
	if err := tmp.Inverse(a); err != nil {
		return nil, err
	}
	return tmp, nil
}

// 线上 Inference
func LinUCBInference(A [][]float64, b []float64, features []float64) (s float64, err error) {

	// 基础参数
	a := 0.1

	featuresMatrix := mat.NewVecDense(len(features), features)
	fmt.Printf("featuresMatrix = %v\n\n", mat.Formatted(featuresMatrix, mat.Prefix("       ")))

	AMatrixInverse, err1 := Inverse(GetMatrix(A))
	if err1 != nil {
		return 0, err1
	}
	fmt.Printf("p^-1 = %v\n\n", mat.Formatted(AMatrixInverse, mat.Prefix("       ")))

	vector := mat.NewVecDense(len(b), b)
	fmt.Printf("vector = %v\n\n", mat.Formatted(vector, mat.Prefix("       ")))

	// 基础计算

	tmpDotLeft := Multi(Multi(AMatrixInverse, vector).T(), featuresMatrix)
	fmt.Printf("dotResult = %v\n\n", mat.Formatted(tmpDotLeft, mat.Prefix("       ")))

	tmpDotRight := Multi(Multi(featuresMatrix.T(), AMatrixInverse), featuresMatrix)
	fmt.Printf("tmpDotRight1 = %v\n\n", mat.Formatted(tmpDotRight, mat.Prefix("       ")))

	result := tmpDotLeft.At(0, 0) + a*math.Sqrt(tmpDotRight.At(0, 0))

	fmt.Printf("最终结果 %0.4v\n\n", result)

	return result, nil
}

//Inference interface
