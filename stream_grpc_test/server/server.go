package main

import (
	"fmt"
	"google.golang.org/grpc"
	"learn/stream_grpc_test/proto"
	"log"
	"net"
	"time"
)

const PORT = ":50052"

type server struct {
	proto.UnimplementedGreeterServer
}

// GetStream 实现了服务器端流式 RPC
func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

// utStream 实现了客户端流式 RPC
func (s *server) utStream(client proto.Greeter_PutStreamClient) error {
	for {
		req, err := client.Recv() // 接收客户端数据
		if err != nil {
			return err
		}
		fmt.Println("Received:", req.GetData())
	}
	return nil
}

// AllStream 实现了双向流式 RPC
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	for {
		req, err := allStr.Recv() // 接收客户端的数据
		if err != nil {
			return err
		}
		fmt.Println("Received:", req.GetData())

		// 发送响应
		err = allStr.Send(&proto.StreamResData{
			Data: fmt.Sprintf("Received: %v", req.GetData()),
		})
		if err != nil {
			return err
		}
	}
}

func main() {
	// 启动 TCP 监听
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// 注册 GreeterServer
	proto.RegisterGreeterServer(s, &server{})

	// 启动 gRPC 服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
