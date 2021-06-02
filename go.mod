module awesomeProject

go 1.16

// 1. 初始化init
// go mod help
// 不要用go get ，会吧包全部下到gopath（c盘）
// 直接复制到本项目下面
require (
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.3.0
)
