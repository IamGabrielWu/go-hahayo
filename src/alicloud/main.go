package main

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func main() {
	// 创建ecsClient实例
	ecsClient, err := ecs.NewClientWithAccessKey(
		"cn-qingdao", // 地域ID
		"xxx",        // 您的Access Key ID
		"xxx")        // 您的Access Key Secret
	if err != nil {
		// 异常处理
		panic(err)
	}
	// 创建API请求并设置参数
	request := ecs.CreateDescribeInstancesRequest()
	// 等价于 request.PageSize = "10"
	request.PageSize = requests.NewInteger(10)
	// 发起请求并处理异常
	response, err := ecsClient.DescribeInstances(request)
	if err != nil {
		// 异常处理
		panic(err)
	}
	fmt.Println(response)
}
