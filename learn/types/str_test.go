package cache_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Test_createStr(t *testing.T) {
	file, err := os.Open("./file")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Println(string(content))

	//This argument must have a pointer type/
	// m := make(map[string][]ProjectDataSourceSdkConfig)
	var m DataSourceSdkConfigs



	json.Unmarshal(content, &m)

	fmt.Println(m)

	//for _, config := range m {
	//
	//	fmt.Println(config)
	//	for _, sdkConfig := range config {
	//		// 	fmt.Println(sdkConfig)
	//
	//		if sdkConfig.Name == "zeus" {
	//			fmt.Println("------")
	//			fmt.Println(sdkConfig.Name)
	//			fmt.Println(sdkConfig.DataSource)
	//			fmt.Println("------")
	//		}
	//	}
	//
	//}
	//fmt.Println(m)

}
