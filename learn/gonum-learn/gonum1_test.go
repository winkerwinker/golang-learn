package gonum_test

//import (
//	"fmt"
//	"gonum.org/v1/gonum/mat"
//	"reflect"
//	"testing"
//)
//
////多维度数组转 matrix
//func GetMatrix(input [][]float64) *mat.Dense {
//	tmpInput := make([]float64, 0, len(input)*len(input[0]))
//
//	for i := 0; i < len(input); i++ {
//		tmpInput = append(tmpInput, input[i]...)
//	}
//	return mat.NewDense(len(input), len(input[0]), tmpInput)
//
//}
//
////多维度数组转 matrix
//func GetMatrixZero(input [][]float64) *mat.Dense {
//	return mat.NewDense(len(input), len(input[0]), nil)
//
//}
//
//func TestLinUCBInference(t *testing.T) {
//	type args struct {
//		A        [][]float64
//		b        []float64
//		features []float64
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantS   fmt.Formatter
//		wantErr bool
//	}{
//		{name: "test1",
//			args:    args{A: [][]float64{{1, 12, 8}, {0, 4, 5}, {0, 0, 9}}, b: []float64{0.1, 0.6, 0.25}, features: []float64{0, 1, 1}},
//			wantS:   nil,
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			gotS, err := LinUCBInference(tt.args.A, tt.args.b, tt.args.features)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("LinUCBInference() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotS, tt.wantS) {
//				t.Errorf("LinUCBInference() gotS = %v, want %v", gotS, tt.wantS)
//			}
//		})
//	}
//}
//
//// 包装乘法方法
//func Multi(a, b mat.Matrix) (*mat.Dense, float64) {
//
//	// 计算点积
//	if GetMatrixRow(a) == 1 && GetMatrixRow(b) == 1 && GetMatrixCol(a) == GetMatrixCol(b) {
//		f := dot(a, b)
//		return nil, f
//	}
//
//	tmpDense := mat.NewDense(GetMatrixRow(a), GetMatrixCol(b), nil)
//	tmpDense.Mul(a, b)
//	return tmpDense, 0
//}
//
//func dot(a, b mat.Matrix) (result float64) {
//	// mat.Matrix is mat.Transpose, not *mat.Dense [recovered]
//	//denseA := a.(*mat.Dense)
//	//denseB := a.(*mat.Dense)
//	// desne 转 vector
//	_, c := a.Dims()
//
//	for i := 0; i < c; i++ {
//		result += a.At(0, i) * b.At(0, i)
//	}
//	return result
//}
//
//func GetMatrixRow(b mat.Matrix) int {
//
//	i, _ := b.Dims()
//	return i
//
//}
//
//func GetMatrixCol(b mat.Matrix) int {
//	_, j := b.Dims()
//	return j
//}
//
//// 包装乘法方法
//func Inverse(a mat.Matrix) (*mat.Dense, error) {
//	i, j := a.Dims()
//	tmp := mat.NewDense(i, j, nil)
//	// AMatrixInverse := GetMatrixZero(A)
//	if err := tmp.Inverse(a); err != nil {
//		return nil, err
//	}
//	return tmp, nil
//}
//
//
//// 线上 Inference
//func LinUCBInference(A [][]float64, b []float64, features []float64) (s fmt.Formatter, err error) {
//
//	// a := 0.1
//
//	featuresMatrix := mat.NewVecDense( len(features), features)
//
//	featuresMatrixT := featuresMatrix.T()
//	fmt.Printf("featuresMatrix = %v\n\n", mat.Formatted(featuresMatrix, mat.Prefix("       ")))
//
//	AMatrixInverse, err1 := Inverse(GetMatrix(A))
//	if err1 != nil {
//		return nil, err1
//	}
//
//	fmt.Printf("p^-1 = %v\n\n", mat.Formatted(AMatrixInverse, mat.Prefix("       ")))
//
//	vector := mat.NewVecDense(len(b), b)
//	fmt.Printf("vector = %v\n\n", mat.Formatted(vector, mat.Prefix("       ")))
//
//	theta, _ := Multi(AMatrixInverse, vector)
//	fmt.Printf("theta = %v\n\n", mat.Formatted(theta, mat.Prefix("       ")))
//	thetaT := theta.T()
//
//	fmt.Printf("thetaT = %v\n\n", mat.Formatted(thetaT, mat.Prefix("       ")))
//
//	tmpDotLeft, _ := Multi(thetaT, featuresMatrix)
//	fmt.Printf("dotResult = %v\n\n", mat.Formatted(tmpDotLeft, mat.Prefix("       ")))
//
//	tmpDotRight1, _ := Multi(featuresMatrixT, AMatrixInverse)
//	fmt.Printf("tmpDotRight1 = %v\n\n", mat.Formatted(tmpDotRight1, mat.Prefix("       ")))
//
//	tmpDotRight2, _ := Multi(tmpDotRight1, featuresMatrix)
//	fmt.Printf("tmpDotRight1 = %v\n\n", mat.Formatted(tmpDotRight2, mat.Prefix("       ")))
//	//
//	////Apply()函数可以让用户自定义任何函数 来操作矩阵
//	//sqrt := mat.NewDense(0, 0, nil)
//	//multiResult := mat.NewDense(0, 0, nil)
//	//result := mat.NewDense(0, 0, nil)
//	//// 定义平方根函数
//	//sqrtFunc := func(_, _ int, v float64) float64 { return math.Sqrt(v) }
//	//multiFunc := func(_, _ int, v float64) float64 { return v * a }
//	//
//	//sqrt.Apply(sqrtFunc, &tmpDotRight1)
//	//fh := mat.Formatted(sqrt, mat.Prefix("         "))
//	//fmt.Printf("矩阵A的开平方 %0.4v\n\n", fh)
//	//
//	//multiResult.Apply(multiFunc, sqrt)
//	//result.Add(multiResult, tmpDotLeft)
//	//
//	//formatted := mat.Formatted(result, mat.Prefix("         "))
//	//fmt.Printf("最终结果 %0.4v\n\n", formatted)
//
//	return nil, nil
//}
