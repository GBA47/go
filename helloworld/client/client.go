package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//1.建立链接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("链接失败")
	}
	// var reply *string = new(string)
	var reply string
	err = client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
