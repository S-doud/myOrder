package main

import (
	"context"
	pb "github.com/S-doud/myProto/order"
	"orderServer/logic"
)

type OrderServiceImpl struct {
	pb.UnimplementedOrderServiceServer
}

func (s *OrderServiceImpl) QueryOrderInfo(ctx context.Context, request *pb.QueryOrderRequest) (*pb.QueryOrderReply, error) {

	err := logic.QueryOrder(ctx, request)
	if err != nil {
		panic(err)
	}

	rsp := &pb.QueryOrderReply{
		ResultCode: "0",
		ResultInfo: "ok",
	}
	return rsp, nil
}

func (s *OrderServiceImpl) mustEmbedUnimplementedOrderServiceServer() {

}
