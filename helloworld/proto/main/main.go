package main

import (
	"fmt"
	// 使用别名避免与 proto 包冲突
	protoPkg "github.com/golang/protobuf/proto"
	// 导入你自己定义的 helloworld/proto 包
	"learn/helloworld/proto"
)

func main() {
	// 创建 HelloRequest 请求
	req := proto.HelloRequest{
		Name: "bobby",
	}

	// 使用 protoPkg（即 github.com/golang/protobuf/proto）来进行 Marshal 操作
	rsp, err := protoPkg.Marshal(&req)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	// 打印序列化后的二进制数据
	fmt.Println(rsp)
}
