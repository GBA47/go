package main

import (
	"fmt"
	"google.golang.org/grpc"
	"learn/stream_grpc_test/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewStreamServiceClient(conn)
	res, _ := c.GetStream(context.BAckground(), &proto.StreamReqData{Data: "默克"})
	for {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
		
	}
}
