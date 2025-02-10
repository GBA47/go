package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main(){
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	def conn.Close()
	c := proto.NewGreeter(conn)
	r, err := c.SayHello(context.Background(), proto.HelloRequesr{Name:"bobby"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}