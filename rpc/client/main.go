package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponseData struct {
	Data int `json:"data"`
}

// Add 函数调用远程服务并返回计算结果
func Add(a, b int) int {
	// 构建请求URL，带上参数
	url := fmt.Sprintf("http://127.0.0.1:8000/add?a=%d&b=%d", a, b)

	// 发送 GET 请求
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return 0
	}
	defer res.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0
	}

	// 打印响应体内容
	fmt.Println("Response Body:", string(body))

	// 解析JSON数据
	var rsaData ResponseData
	err = json.Unmarshal(body, &rsaData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return 0
	}

	return rsaData.Data
}

func main() {
	// 调用 Add 函数并打印结果
	fmt.Println("Result:", Add(1, 2))
}
