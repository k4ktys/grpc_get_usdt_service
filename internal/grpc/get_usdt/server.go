package getusdt

import (
	"context"
	"grpc_get_usdt_service/internal/domain/models"
	pb "grpc_get_usdt_service/protos/gen/go/get_usdt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetUsdtServicer interface {
	GetRates(ctx context.Context, market string) (models.UsdtTrade, error)
}

type serverAPI struct {
	pb.UnimplementedGetUsdtServer
	get_usdt_service GetUsdtServicer
}

func Register(gRPC *grpc.Server, service GetUsdtServicer) {
	pb.RegisterGetUsdtServer(gRPC, &serverAPI{get_usdt_service: service})
}

func (s *serverAPI) GetRates(ctx context.Context, req *pb.GetRatesRequest) (*pb.GetRatesResponse, error) {
	if req.Market == "" {
		return nil, status.Error(codes.InvalidArgument, "market id is required")
	}

	trade, err := s.get_usdt_service.GetRates(ctx, req.Market)

	if err != nil {
		return nil, status.Error(codes.Internal, "service layer error: "+err.Error())
	}

	return &pb.GetRatesResponse{
		Timestamp: trade.Timestamp,
		Ask: &pb.GetRatesResponse_Candle{
			Price:  trade.Asks[0].Price,
			Volume: trade.Asks[0].Volume,
			Amount: trade.Asks[0].Amount,
			Factor: trade.Asks[0].Factor,
			Type:   trade.Asks[0].Type,
		},
		Bid: &pb.GetRatesResponse_Candle{
			Price:  trade.Bids[0].Price,
			Volume: trade.Bids[0].Volume,
			Amount: trade.Bids[0].Amount,
			Factor: trade.Bids[0].Factor,
			Type:   trade.Bids[0].Type,
		},
	}, nil
}
