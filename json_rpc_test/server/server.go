package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply的值
	*reply = "Hello, " + request
	return nil
}
func main() {
	//1.实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	//2.注册处理逻辑handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	//3.启动服务
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
