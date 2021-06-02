package adapter

import "fmt"

// 创建云主机
type ICreateSver interface {
	//表示都是 float64
	CreateServer(cpu, mem float64) error
}

type AWSClient struct {
}

func (A AWSClient) CreateServer(cpu, mem float64) error {
	fmt.Println("run")
	return nil
}

type AliyunClient struct {
}

// 使用的是int 需要套接一层调用
func (A AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Println("run aliyun")
	return nil
}

type AliyunClientAdapter struct {
	Client AliyunClient
}

func (A AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	A.Client.CreateServer(int(cpu), int(mem))
	return nil
}
