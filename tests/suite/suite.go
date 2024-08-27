package suite

import (
	"context"
	"grpc_get_usdt_service/internal/config"
	pb "grpc_get_usdt_service/protos/gen/go/get_usdt"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T
	Cfg           *config.Config
	GetUsdtClient pb.GetUsdtClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg, err := config.NewConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Second*60)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.NewClient("localhost:44044", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	return ctx, &Suite{
		T:             t,
		Cfg:           cfg,
		GetUsdtClient: pb.NewGetUsdtClient(cc),
	}
}
