package getusdt

import (
	"context"
	pb "grpc_get_usdt_service/protos/gen/go/get_usdt"

	"google.golang.org/grpc"
)

type serverAPI struct {
	pb.UnimplementedGetUsdtServer
}

func Register(gRPC *grpc.Server) {
	pb.RegisterGetUsdtServer(gRPC, &serverAPI{})
}

func (s *serverAPI) GetRates(ctx context.Context, req *pb.GetRatesRequest) (*pb.GetRatesResponse, error) {
	panic("unimpls")
}
