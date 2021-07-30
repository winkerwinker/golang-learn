package cache_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

//
//import (
//	"encoding/json"
//	"fmt"
//	"math/rand"
//	"strconv"
//	"testing"
//)
//
func Test_createStr8(t *testing.T) {
	features := InputFeature{
		FeaturesType:   "float32",
		OperationName1: "serving_default_leaf_idx_input",
		OperationName2: "StatefulPartitionedCall",
		Length:         339,
	}

	config := DeepModelRawConfig{
		InputFeature: features,
		Params:       "serve",
	}

	fmt.Println(config)
	marshal, _ := json.Marshal(config)
	fmt.Println(string(marshal))
}

// 深度模型在apollo的配置
type DeepModelRawConfig struct {
	InputFeature InputFeature
	Params       string
	//HeatTime     int
}

// 对入参的定义
type InputFeature struct {
	FeaturesType   string
	Length         int
	OperationName1 string
	OperationName2 string
	//Max          string
	//Min          string
}

// 深度模型在apollo的配置
//type DeepModelRawConfig struct {
//	InputFeatures InputFeatures `json:inputFeatures`
//	Params        string        `json:params       `
//	HeatTime      int           `json:heatTime     `
//}
//
//// 对入参的定义
//type InputFeatures struct {
//	FeaturesType string `json: 	featuresType`
//	Max          string `json: 	max         `
//	Min          string `json: 	min         `
//	Length       int    `json: 	length      `
//}

//// 数据转换三种方案
///*
//string(xxx)
//xxx.(string)
//(string)xxx
//
//
//*/
//
////xgb Fm 预热
//func preHeatXgbFMDeepDemo(modelName string, model *tf.SavedModel) {
//	TreeNums := 300
//	feature := randInt64Leaves(TreeNums)
//	FMModelServer(model, feature)
//}
//func preHeatXgbFMv2DeepDemo(modelName string, model *tf.SavedModel) {
//	TreeNums := 316
//	feature := randInt64Leaves(TreeNums)
//
//	FMModelServer(model, feature)
//}
//
////随机生成xgb叶子节点 叶子节点编号0-62
//func randInt64Leaves(length int) [][]int64 {
//	var min int = 0
//	var max int = 31
//	res := make([][]int64, 1)
//	for i := range res {
//		res[i] = make([]int64, length)
//		for j := range res[i] {
//			res[i][j] = int64(min + rand.Intn(max-min+1))
//		}
//	}
//	return res
//}
//
//// todo
//func preHeatXgbFMv3DeepDemo(modelName string, model *tf.SavedModel) {
//	feature := randFloatsArr(380)
//	FMModelServer(model, feature)
//}
//
//
//
//// todo
//func randsArrBase(inputFeature InputFeature) (interface{}, error) {
//	featuresType := inputFeature.FeaturesType
//	length := inputFeature.Length
//	max := inputFeature.Max
//	min := inputFeature.Min
//
//	switch featuresType {
//	case "int":
//		min, err1 := strconv.Atoi(min)
//		max, err2 := strconv.Atoi(max)
//		if err1 != nil || err2 != nil {
//			return nil, errors.New("cannot convert min/max to  featuresType")
//		}
//		// todo 确定一组就可以吗？
//		res := make([][]int, 1)
//		res[0]=make()
//		for i := range res {
//			res[i] = make([]int, length)
//			for j := range res[i] {
//				res[i][j] = min + rand.Int()*(max-min)
//			}
//		}
//		return res, nil
//	case "int64":
//		min, err1 := strconv.ParseInt(min, 10, 64)
//		max, err2 := strconv.ParseInt(max, 10, 64)
//		if err1 != nil || err2 != nil {
//			return nil, errors.New("cannot convert min/max to  featuresType")
//		}
//		// todo 确定一组就可以吗？
//		res := make([][]int64, 1)
//		for i := range res {
//			res[i] = make([]int64, length)
//			for j := range res[i] {
//				res[i][j] = min + rand.Int63()*(max-min)
//			}
//		}
//		return res, nil
//	case "float32":
//		min, err1 := strconv.ParseFloat(min, 32)
//		max, err2 := strconv.ParseFloat(max, 32)
//		if err1 != nil || err2 != nil {
//			return nil, errors.New("cannot convert min/max to  featuresType")
//		}
//		res := make([][]float32, 1)
//		for i := range res {
//			res[i] = make([]float32, length)
//			for j := range res[i] {
//				res[i][j] = float32(min) + rand.Float32()*(float32(max)-float32(min))
//			}
//		}
//		return res, nil
//	case "float64":
//		min, err1 := strconv.ParseFloat(min, 64)
//		max, err2 := strconv.ParseFloat(max, 64)
//		if err1 != nil || err2 != nil {
//			return nil, errors.New("cannot convert min/max to  featuresType")
//		}
//		res := make([][]float64, 1)
//		for i := range res {
//			res[i] = make([]float64, length)
//			for j := range res[i] {
//				res[i][j] = min + rand.Float64()*(max-min)
//			}
//		}
//		return res, nil
//	default:
//		return nil, errors.New("cannot find featuresType randsArrBase function")
//	}
//
//}
