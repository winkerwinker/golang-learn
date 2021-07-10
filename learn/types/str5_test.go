package cache_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_json(t *testing.T) {

	//loaderConf := &LoaderConf{
	//	Index: 0,
	//	Fusion: FusionConf{
	//		Ip:     "10.85.129.62",
	//		Port:   4600,
	//		Key:    "cxyx_xgb_goodfeature_v1_protobuf_encoded:<INPUTGOODSID>",
	//		Prefix: "test"},
	//	Opr:        "",
	//	ResultType: "",
	//	Apollo: ApolloConf{
	//		Namespace: "",
	//		ConfName:  "",
	//	},
	//}
	//marshal, _ := json.Marshal(loaderConf)
	//fmt.Println(string(marshal))

	s := "{\"hello\":\"10.85.129.62\",\"Ip\":\"10.85.129.62\",\"Port\":4600,\"Prefix\":\"test\",\"Key\":\"cxyx_xgb_goodfeature_v1_protobuf_encoded:\\u003cINPUTGOODSID\\u003e\"}"
	conf1 := &FusionConf{}
	err := json.Unmarshal([]byte(s), conf1)
	fmt.Println(err)
	fmt.Println(conf1)
	fmt.Println("========== ")
	conf := FusionConf{
		Ip:     "10.85.129.62",
		Port:   4600,
		Key:    "cxyx_xgb_goodfeature_v1_protobuf_encoded:<INPUTGOODSID>",
		Prefix: "test"}

	marshal, _ := json.Marshal(conf)
	fmt.Println(string(marshal))
}

type LoaderConf struct {
	// engine 中的标识
	Index      int
	Fusion     FusionConf
	Apollo     ApolloConf
	Opr        string
	ResultType string
}

type FusionConf struct {
	Ip     string
	Port   int    //todo ip port 替换成fusion 标识T1 T2
	Prefix string //todo 看如何解析
	Key    string
}
type ApolloConf struct {
	Namespace string
	ConfName  string
}

//
//{
//"defaultSorter": {
//"model": {
//"modelCode": "DetailPageMiddleXGBSorterV1",
//"loadFrom": "remote",
//"filePath": ""
//},
//"data": [
//{
//"dataToken": "",
//"source": {
//"featureType": "T4",
//"HasCache": false,
//"CacheName": "",
//"CacheSchema": "",
//"FusionSchema": "cxyx_recommender_user_features_xgb_v1:<USERID>",
//"FusionAddr": "100.69.239.198:3140",
//"FusionField": "features"
//},
//"concat": {
//"ConcatKey": "<USERID>",
//"Order": 0
//},
//"size": 292
//},
//{
//"dataToken": "",
//"source": {
//"featureType": "T9",
//"HasCache": true,
//"CacheName": "cache_1",
//"CacheSchema": "a_goodfeature_v1:<INPUTGOODSID>",
//"FusionSchema": "cxyx_xgb_goodfeature_v1_protobuf_encoded:<INPUTGOODSID>",
//"FusionAddr": "10.85.129.62:4600",
//"FusionField": "features"
//},
//"concat": {
//"ConcatKey": "<INPUTGOODSID>",
//"Order": 1
//},
//"size": 116
//},
//{
//"dataToken": "",
//"source": {
//"featureType": "T9",
//"HasCache": true,
//"CacheName": "cache_1",
//"CacheSchema": "a_goodfeature_v1:<GOODSID>",
//"FusionSchema": "cxyx_xgb_goodfeature_v1_protobuf_encoded:<GOODSID>",
//"FusionAddr": "10.85.129.62:4600",
//"FusionField": "features"
//},
//"concat": {
//"ConcatKey": "<GOODSID>",
//"Order": 2
//},
//"size": 116
//},
//{
//"dataToken": "",
//"source": {
//"featureType": "T3",
//"HasCache": false,
//"CacheName": "",
//"CacheSchema": "",
//"FusionSchema": "youxuan_user_product_paired:<USERID>",
//"FusionAddr": "100.69.239.199:3250",
//"FusionField": "<GOODSID>"
//},
//"concat": {
//"ConcatKey": "<GOODSID>",
//"Order": 3
//},
//"size": 68
//}
//]
//},
//"isGroup": false
//}
