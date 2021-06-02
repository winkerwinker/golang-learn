package map_test

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
)

func TestMap4(t *testing.T) {

	//os.RemoveAll("/tmp/")

	err := os.RemoveAll("/Users/didi/go/src/awesomeProject/learn/map/test1")

	fmt.Println(err)
	fmt.Println("--------------")
	strs := []string{"2620369141", "1820369141", "1620359341", "1720359341"}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i] > strs[j]
	})

	fmt.Println(strs)

	tempFile := "XGBSorterV2_v37.model_"
	// XGBSorterV2_lgb_v37.model  ->  modelCode= XGBSorterV2，Version=model
	index := strings.LastIndex(tempFile, "_")
	modelCode := tempFile[:index]
	modelVersion := tempFile[index+1:]

	// 适配文件和文件夹
	if strings.Contains(modelVersion, ".") {
		modelVersion = strings.Split(tempFile[index+1:], ".")[0]
	}
	lastIndex := strings.LastIndex(modelCode, "_")
	fmt.Println(lastIndex)
	modelType := modelCode[lastIndex+1:]
	modelCode = modelCode[:lastIndex]

	fmt.Println(modelType)
	fmt.Println(modelCode)
	fmt.Println(modelVersion)
}

// map中找不到的key

//model-engine.log.1 model-engine.log.10 model-engine.log.2 model-engine.log.3 model-engine.log.4 model-engine.log.5 model-engine.log.6 model-engine.log.7 model-engine.log.8
//
//
//hadoop fs -put model_deep_v2  /user/prod_cxyx_growth_recom/cxyx_growth_recom/hdp/xieaichen
//hadoop fs -put model_deep_v1  /user/prod_cxyx_growth_recom/cxyx_growth_recom/hdp/xieaichen
//hadoop fs -put model_deep_v3  /user/prod_cxyx_growth_recom/cxyx_growth_recom/hdp/xieaichen
//hadoop fs -put model_deep_v4  /user/prod_cxyx_growth_recom/cxyx_growth_recom/hdp/xieaichen
//hadoop fs -put model_deep_v6  /user/prod_cxyx_growth_recom/cxyx_growth_recom/hdp/xieaichen
//hadoop fs -put model_deep_v7  /user/prod_cxyx_growth_recom/cxyx_growth_recom/hdp/xieaichen