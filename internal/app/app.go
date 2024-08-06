package app

import (
	grpcapp "grpc_get_usdt_service/internal/app/grpc"

	"go.uber.org/zap"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *zap.Logger, grpcPort int) *App {
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
