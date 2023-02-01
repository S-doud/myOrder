package main

import (
	"context"
	"fmt"
	pb "github.com/S-doud/myProto/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	req := &pb.QueryOrderRequest{
		OrderId: "1241254",
	}

	orderServerClient := pb.NewOrderServiceClient(conn)

	rsp, err := orderServerClient.QueryOrderInfo(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.ResultCode)
	fmt.Println(rsp.ResultInfo)

}
