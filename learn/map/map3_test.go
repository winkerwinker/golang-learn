package map_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

type demo struct {
	str  string `json:"str"`
	sex string `json:"sex"`
}

//
//func randFloatsArr(length int) interface{} {
//
//	tempFile := "XGBSorterV2_lgb_v37.mode"
//	// XGBSorterV2_lgb_v37.model  ->  modelCode= XGBSorterV2，Version=model
//	index := strings.LastIndex(tempFile, "_")
//	modelCode := tempFile[:index]
//	modelVersion := tempFile[index+1:]
//
//	// 适配文件和文件夹
//	if strings.Contains(modelVersion, ".") {
//		modelVersion = strings.Split(tempFile[index+1:], ".")[0]
//	}
//	lastIndex := strings.LastIndex(modelCode, "_")
//	modelType := modelCode[lastIndex+1:]
//	modelCode = modelCode[:lastIndex]
//
//
//	fmt.Println(modelType)
//	fmt.Println(modelCode)
//	fmt.Println(modelVersion)
//
//	//var min float64 = 0.0001
//	//var max float64 = 0.0002
//	//res := make([][]float64, 1)
//	//for i := range res {
//	//	res[i] = make([]float64, length)
//	//	for j := range res[i] {
//	//		res[i][j] = min + rand.Float64()*(max-min)
//	//	}
//	//}
//	//return res
//}

//[[0.00016046602879796196 0.00019405090880450125 0.00016645600532184905

type ModelInstanceWrapper struct {
	bizCode        string      `json:"bizCode"`        //编码
	modelCode      string      `json:"modelCode"`      //编码
	modelInstances interface{} `json:"modelInstances"` //真正的模型实例
	createTime     time.Time   `json:"createTime"`     //时间戳
	timeStamp      string      `json:"timeStamp"`      //加载时间
	modelVersion   string      `json:"modelVersion"`   //版本
}

func TestMap3(t *testing.T) {

	wrapper := ModelInstanceWrapper{bizCode: "jhh", modelVersion: "888"}
	marshal, _ := json.Marshal(wrapper)
	fmt.Println(string(marshal))
	fmt.Println("----")

	d := demo{str: "1111", sex: "nv"}
	//dd := demo{str: "111122", serx: "nv"}
	marshal1, _ := json.Marshal(d)
	fmt.Println(string(marshal1))


	fmt.Println("----")
	str := "[{\"price\":\"0\",\"goods_id\":\"20972\"},{\"price\":\"0\",\"goods_id\":\"23216\"},{\"price\":\"0\",\"goods_id\":\"25032\"},{\"price\":\"0\",\"goods_id\":\"31146\"},{\"price\":\"0\",\"goods_id\":\"39744\"},{\"price\":\"0\",\"goods_id\":\"8793\"},{\"price\":\"0\",\"goods_id\":\"35196\"},{\"price\":\"0\",\"goods_id\":\"23570\"},{\"price\":\"0\",\"goods_id\":\"38022\"},{\"price\":\"0\",\"goods_id\":\"20162\"},{\"price\":\"0\",\"goods_id\":\"11881\"},{\"price\":\"0\",\"goods_id\":\"43598\"},{\"price\":\"0\",\"goods_id\":\"15227\"},{\"price\":\"0\",\"goods_id\":\"9003\"},{\"price\":\"0\",\"goods_id\":\"22856\"},{\"price\":\"0\",\"goods_id\":\"23120\"},{\"price\":\"0\",\"goods_id\":\"43600\"},{\"price\":\"0\",\"goods_id\":\"11871\"},{\"price\":\"0\",\"goods_id\":\"25476\"},{\"price\":\"0\",\"goods_id\":\"8739\"},{\"price\":\"0\",\"goods_id\":\"23732\"},{\"price\":\"0\",\"goods_id\":\"23838\"},{\"price\":\"0\",\"goods_id\":\"11411\"},{\"price\":\"0\",\"goods_id\":\"23704\"},{\"price\":\"0\",\"goods_id\":\"22434\"},{\"price\":\"0\",\"goods_id\":\"31486\"},{\"price\":\"0\",\"goods_id\":\"18336\"},{\"price\":\"0\",\"goods_id\":\"23126\"},{\"price\":\"0\",\"goods_id\":\"19266\"},{\"price\":\"0\",\"goods_id\":\"13733\"},{\"price\":\"0\",\"goods_id\":\"23720\"},{\"price\":\"0\",\"goods_id\":\"26680\"},{\"price\":\"0\",\"goods_id\":\"34738\"},{\"price\":\"0\",\"goods_id\":\"38858\"},{\"price\":\"0\",\"goods_id\":\"19576\"},{\"price\":\"0\",\"goods_id\":\"19012\"},{\"price\":\"0\",\"goods_id\":\"44220\"},{\"price\":\"0\",\"goods_id\":\"22432\"},{\"price\":\"0\",\"goods_id\":\"23812\"},{\"price\":\"0\",\"goods_id\":\"23004\"},{\"price\":\"0\",\"goods_id\":\"37838\"},{\"price\":\"0\",\"goods_id\":\"44676\"},{\"price\":\"0\",\"goods_id\":\"45678\"},{\"price\":\"0\",\"goods_id\":\"40100\"},{\"price\":\"0\",\"goods_id\":\"31206\"},{\"price\":\"0\",\"goods_id\":\"37020\"}]"
	m := make(map[string]string, 10)
	_ = json.Unmarshal([]byte(str), &m)

	var features64 []float64
	for _, f64 := range features64 {
		fmt.Println(f64)
	}

	fmt.Println("xxx =-=====")
	//arr := randFloatsArr(100)
	//result := arr.([][]float32)
	//fmt.Println(result)

	////开箱即用
	//var sm sync.Map
	//d := demo{str: "1111", serx: "nv"}
	//dd := demo{str: "111122", serx: "nv"}
	//marshal, _ := json.Marshal(d)
	//fmt.Println(string(marshal))
	//
	//var demos []demo
	//demos = append(demos, d)
	//demos = append(demos, dd)
	//bytes, _ := json.Marshal(demos)
	//fmt.Println(string(bytes))
	//
	//fmt.Println("-------------------------")
	//fmt.Println("-------------------------")
	//
	////store 方法,添加元素
	//sm.Store(1, "a")
	////Load 方法，获得value
	//if v, ok := sm.Load(1); ok {
	//	fmt.Println(v)
	//}
	////LoadOrStore方法，获取或者保存
	////参数是一对key：value，如果该key存在且没有被标记删除则返回原先的value（不更新）和true；不存在则store，返回该value 和false
	//if vv, ok := sm.LoadOrStore(1, "c"); ok {
	//	fmt.Println(vv)
	//}
	//if vv, ok := sm.LoadOrStore(2, "c"); !ok {
	//	fmt.Println(vv)
	//}
	////遍历该map，参数是个函数，该函数参的两个参数是遍历获得的key和value，返回一个bool值，当返回false时，遍历立刻结束。
	//fmt.Println("-------------------------")
	//fmt.Println(sm)
	//marshal, _ := json.Marshal(sm)
	//fmt.Println(string(marshal))
	//fmt.Println("-------------------------")
	//
	//sm.Range(func(k, v interface{}) bool {
	//	fmt.Print(k)
	//	fmt.Print(":")
	//	fmt.Print(v)
	//	fmt.Println()
	//	return true
	//})

	fmt.Println(strings.HasSuffix("test", ""))
	//fmt.Println(strings.HasSuffix("test", nil))
	//fmt.Println(strings.HasSuffix("test", ""))
	//fmt.Println(strings.HasSuffix("test", ""))

	//d := demo{str: "okok"}
	//
	//ints := []int{1, 2, 3, 4}
	//fmt.Println(reflect.TypeOf(ints).Comparable())
	//fmt.Println("-----------")
	//
	//of := reflect.TypeOf(d)
	//off := reflect.TypeOf(of)
	//// comparable := off.Comparable()
	//comparable := off.Comparable()
	//fmt.Println(comparable)
	//
	//var wg sync.WaitGroup
	//wg.Add(3)
	//go func() {
	//	defer wg.Done()
	//
	//	fmt.Println("hello1")
	//	time.Sleep(30 * time.Second)
	//	fmt.Println("qqq")
	//}()
	//go func() {
	//	defer wg.Done()
	//
	//	fmt.Println("hello1")
	//	time.Sleep(30 * time.Second)
	//	fmt.Println("zz")
	//}()
	//go func() {
	//	defer wg.Done()
	//
	//	fmt.Println("hello1")
	//	time.Sleep(30 * time.Second)
	//	fmt.Println("qttt")
	//}()
	//
	//wg.Wait()
	//fmt.Println("bye")
}
