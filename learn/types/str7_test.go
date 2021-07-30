package cache_sdk

//
//import (
//	"fmt"
//	"github.com/sirupsen/logrus"
//	"math/rand"
//	"reflect"
//	"testing"
//)
//
//func Test_test7(t *testing.T) {
//	//before := make([]int, 1)
//	////outputFeatures := make([]int, 10)
//	////outputFeatures = nil
//	//var outputFeatures []int
//	//before = append(before, outputFeatures...)
//}
//
//
//
//package bifrost_model_loader
//
//import (
//"fmt"
//"math/rand"
//"middle_server/logger"
//"middle_server/niubipkg/sync"
//"reflect"
//
//"github.com/sirupsen/logrus"
//tf "github.com/tensorflow/tensorflow/tensorflow/go"
//)
//
//var deepModelConfigMap *sync.ConcurrentMap
//
//type DeepModelConfig struct {
//	tags      []string
//	preHeat   func(modelName string, model *tf.SavedModel)
//	heatTimes int
//}
//
//func init() {
//	deepModelConfigMap, _ = sync.NewConcurrentMap(reflect.TypeOf("placeholder"))
//
//	registerDeepModel()
//}
//
//// 深度 上线前注册tag,注册预热方法,注册预热方法加载次数
//func registerDeepModel() {
//	config := &DeepModelConfig{tags: []string{"serve"}, preHeat: preHeatDemo, heatTimes: 5}
//	xgbFMDeepConfig := &DeepModelConfig{tags: []string{"serve"}, preHeat: preHeatXgbFMDeepDemo, heatTimes: 2}
//	xgbFMv2DeepConfig := &DeepModelConfig{tags: []string{"serve"}, preHeat: preHeatXgbFMv2DeepDemo, heatTimes: 2}
//	xgbFMv3DeepConfig := &DeepModelConfig{tags: []string{"serve"}, preHeat: preHeatXgbFMv3DeepDemo, heatTimes: 2}
//
//	deepModelConfigMap.Store("model", config)
//	deepModelConfigMap.Store("xgbFM", xgbFMDeepConfig)
//	deepModelConfigMap.Store("xgbFMV2", xgbFMv2DeepConfig)
//	deepModelConfigMap.Store("xgbFMV3", xgbFMv3DeepConfig)
//}
//
//// 深度预热demo
//func preHeatDemo(modelName string, model *tf.SavedModel) {
//	features := randFloatsArr(339)
//	tensor, err := tf.NewTensor(features)
//
//	if err != nil {
//		logger.GetBifrostModelLogger(modelName, "").WithFields(logrus.Fields{"msg": "preheat_create_tensor_failed", "err": err}).Warn(logger.BIFROST_LOAD)
//	}
//	model.Session.Run(
//		map[tf.Output]*tf.Tensor{
//			model.Graph.Operation("serving_default_input_1").Output(0): tensor,
//		},
//		[]tf.Output{
//			model.Graph.Operation("StatefulPartitionedCall").Output(0),
//		},
//		nil,
//	)
//
//}
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
//	FMModelServer(model, feature)
//}
//func preHeatXgbFMv3DeepDemo(modelName string, model *tf.SavedModel) {
//	feature := randFloatsArr(380)
//	FMModelServerFloat32(model, feature)
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
//
//
//func FMModelServerFloat32(model *tf.SavedModel, features [][]float32) ([][]float32, error) {
//	tensor, err := tf.NewTensor(features)
//	if err != nil {
//		fmt.Printf("Error NewTensor: err: %s", err.Error())
//		return nil, err
//	}
//	result, err := model.Session.Run(
//		map[tf.Output]*tf.Tensor{
//			model.Graph.Operation("serving_default_inputs").Output(0): tensor,
//		},
//		[]tf.Output{
//			model.Graph.Operation("StatefulPartitionedCall").Output(0),
//		},
//		nil,
//	)
//	if err != nil {
//		fmt.Printf("Error running the session with input, err: %s  ", err.Error())
//		return nil, err
//	}
//	return result[0].Value().([][]float32), nil
//}
//
//// 特征配置化  -> 某一个用户 uid
//
//// 反射 int float
//// 维度
//// 339
//func randFloatsArr(length int) [][]float32 {
//	var min float32 = 0.0001
//	var max float32 = 0.0002
//	res := make([][]float32, 1)
//	for i := range res {
//		res[i] = make([]float32, length)
//		for j := range res[i] {
//			res[i][j] = min + rand.Float32()*(max-min)
//		}
//	}
//	return res
//}
//
//// 获得深度配置
//func getDeepModelConfigMap(modelCode string) *DeepModelConfig {
//	var (
//		inter      interface{}
//		ok         bool
//		deepConfig *DeepModelConfig
//	)
//
//	if inter, ok = deepModelConfigMap.Load(modelCode); !ok {
//		logger.GetBifrostModelLogger(modelCode, "").WithFields(logrus.Fields{"msg": "cannot find deep model config"}).Warn(logger.BIFROST_LOAD)
//	}
//	deepConfig = inter.(*DeepModelConfig)
//	return deepConfig
//}
//
//func executePreHeat(modelCode string, instance interface{}) *tf.SavedModel {
//
//	deepConfig := getDeepModelConfigMap(modelCode)
//
//	model := instance.(*tf.SavedModel)
//	for i := 0; i < deepConfig.heatTimes; i++ {
//		deepConfig.preHeat(modelCode, instance.(*tf.SavedModel))
//	}
//	return model
//}
