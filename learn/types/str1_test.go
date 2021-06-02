package cache_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_createStr1(t *testing.T) {

	content := `{"delay":180,"fields":"id,new_category_id,price","name":"nissna"}`


	//This argument must have a pointer type/
	// m := make(map[string][]ProjectDataSourceSdkConfig)
	var m DataSourceConfig



	json.Unmarshal([]byte(content), &m)

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
