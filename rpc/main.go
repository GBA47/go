package main

import "fmt"

func Add(a, b int) int {
	total := a + b
	return total
}

type Company struct {
	Name    string
	Address string
}
type Employee struct {
	Name    string
	Company company
}

type PrintResult struct {
	Info string
	Err  error
}

func RpcPrintln(employee Employee) {
	//Rpc中的第二个点 传输协议，数据编码协议
	//http1.x http2.0协议
	//http协议底层也用tcp http主流是1.x 这种协议一旦结果返回，链接就断开
	//1. 直接基于tcp/ucp协议 去封装一层myhttp,没有通用性
	/*客户端
		1.建立链接tcp/http
		2.将employee对象序列化成json字符串-序列化
		3.发送json字符串-调用成功后实际上你接收的是一个二进制的数据
		4.等待服务器发送结果
		5.将服务器返回的结果解析成PrintResult对象 - 反序列化
	服务端
		1.监听网络端口80
	2.读取数据- 二进制的json数据
	3.对数据进行反序列化Employee对象
	4.开始处理业务逻辑
	5.将处理结果PrintResult序列化成json二进制
	6.将服务器返回的数据进行成PrintResult对象 - 反序列化
	*/
}

// 将打印工作放在别的服务器上，应将struct序列化为json
func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Employee{
		Name: "bobby",
		Company: Company{
			Name:    "公司",
			Address: "地址",
		},
	})
	//远程服务应反解
}
