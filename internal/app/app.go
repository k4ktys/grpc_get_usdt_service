package app

import (
	grpcapp "grpc_get_usdt_service/internal/app/grpc"
	"grpc_get_usdt_service/internal/config"
	getusdt "grpc_get_usdt_service/internal/services/get_usdt"
	"grpc_get_usdt_service/internal/storage/postgresql"

	"go.uber.org/zap"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *zap.Logger, cfg *config.Config, grpcPort int) *App {
	storage, err := postgresql.New(cfg)
	if err != nil {
		panic(err)
	}

	service := getusdt.New(log, storage)

	grpcApp := grpcapp.New(log, service, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
