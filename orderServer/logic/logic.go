package logic

import (
	"context"
	"fmt"
	pb "github.com/S-doud/myProto/order"
)

func QueryOrder(ctx context.Context, req *pb.QueryOrderRequest) error {

	fmt.Println("QueryOrder begin")
	fmt.Println(req.OrderId)

	return nil
}
