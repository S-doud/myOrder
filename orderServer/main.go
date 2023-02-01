package main

import (
	"fmt"
	pb "github.com/S-doud/myProto/order"
	"google.golang.org/grpc"
	"net"
)

func main() {
	orderService := &OrderServiceImpl{}
	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, orderService)

	lis, err := net.Listen("tcp", ":9001")
	fmt.Println("server start: tcp 9001")
	if err != nil {
		panic(err)
	}

	server.Serve(lis)
}
