package map_test

import (
	"fmt"
	"path/filepath"
	"sort"
	"testing"
)

/*

指针
Dig101-Go之如何在函数内修改指针指向

*/

type Prediction struct {
	Key    int64
	Value  float64
	Leaves []float64
}
type PredictionResult struct {
	Predictions []*Prediction
}

// 降序排序
func (p PredictionResult) Less(i, j int) bool {
	return p.Predictions[i].Value > p.Predictions[j].Value
}
func (p PredictionResult) Len() int {
	return len(p.Predictions)
}

func (p PredictionResult) Swap(i, j int) {
	p.Predictions[i], p.Predictions[j] = p.Predictions[j], p.Predictions[i]
}

func IF(cond bool, trueVal, falseVal interface{}) interface{} {
	if cond {
		return trueVal
	}
	return falseVal
}

func TestDo8(t *testing.T) {

	fmt.Println(filepath.Join("hhh./", "jjjj", ""))


	errFeatureGid := make([]int64, 0, 10)
	//for i := 0; i < 19; i++ {
	//	errFeatureGid = append(errFeatureGid, 19)
	//}

	//errFeatureGid = append(errFeatureGid, 19)

	printErrFeatureGid := make([]int64, 0, 10)

	fmt.Println(len(errFeatureGid))
	// 有错误特征的商品数据进行埋点
	printErrFeatureGid = IF(
		len(errFeatureGid) > 20,
		errFeatureGid[0:20],
		errFeatureGid).([]int64)

	//if len(errFeatureGid) > 20 {
	//	printErrFeatureGid = errFeatureGid[0:20]
	//} else {
	//	printErrFeatureGid = errFeatureGid
	//}
	fmt.Println(printErrFeatureGid)
	fmt.Println("-------------")

	prediction := PredictionResult{
		Predictions: []*Prediction{
			&Prediction{Key: 1, Value: 2.0},
			&Prediction{Key: 2, Value: 1.0},
			&Prediction{Key: 4, Value: 5.0},
			&Prediction{Key: 3, Value: 88},
			&Prediction{Key: 7, Value: -1},
			&Prediction{Key: 8, Value: -1},
			&Prediction{Key: 7, Value: -1},
		},
	}

	sort.Sort(prediction)

	for i := 0; i < len(prediction.Predictions); i++ {

		fmt.Println(prediction.Predictions[i].Key)
		//fmt.Println(prediction.Predictions[i].Value)
	}

	////m1()
	//fmt.Println("===========")
	//var argType *[]ArgType
	//fmt.Println(argType)
	//modifyPointerArg11(argType)
	//fmt.Println(&argType)
	//fmt.Println(argType)
}

type ArgType struct {
	A string
	b int
}

func modifyPointerArg11(arg *[]ArgType) {
	argType := ArgType{"arg1", 1}
	//  todo 如果要给空指针指向一个新地方会panic
	*arg = []ArgType{argType}
	fmt.Println(arg)
	fmt.Println(*arg)
	//fmt.Println("inside modifyPointerArg1:", arg)
}

// 值传递，值修改了这个复制过来的值并，没有用
func modifyPointerArg1(arg *ArgType) {
	arg = &ArgType{"arg1", 1}
	fmt.Println(arg)
	fmt.Println(*arg)
	//fmt.Println("inside modifyPointerArg1:", arg)
}

// todo
func modifyPointerArg2(arg *ArgType) {
	// todo todo  修改了他指向的值
	*arg = ArgType{"arg2", 2}
	fmt.Println(arg)
	fmt.Println(*arg)
	fmt.Println("inside modifyPointerArg2:", arg)
}

func m1() {
	to := new(string)
	s := "to"
	to = &s
	fmt.Println(to)
	test("ok", to)
	fmt.Println(to)
	fmt.Println(*to)
}

func test(form string, to *string) {
	fmt.Println(to)
	*to = form
	// 不应该把这个带回去吗？
	fmt.Println(to)
	fmt.Println(*to)
}

//0xc000126028
//0xc000110890
//0xc0001108a0
//0xc000126028
//0xc000110890
