package cache_sdk

// 定义json 机构直接解析

//  定义所有类型的数据源， 商品、用户等等
//type DataSourceSdkConfigs struct {
//	dataSourceSdkConfig map[string][]ProjectDataSourceSdkConfig
//}

//todo
 // !!!! 有空再看吧
type DataSourceSdkConfigs struct {
	goodsCache *[]ProjectDataSourceSdkConfig
}


// 不同的项目对数据源的依赖
type ProjectDataSourceSdkConfig struct {
	Name           string
	Delay          int64
	Fields         string
	DataSource *[]DataSourceConfig
}

type DataSourceConfig struct {
	Name   string
	Delay  int64
	Fields string
}
