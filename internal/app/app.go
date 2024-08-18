package app

import (
	grpcapp "grpc_get_usdt_service/internal/app/grpc"
	"grpc_get_usdt_service/internal/config"
	"grpc_get_usdt_service/internal/metrics"
	"grpc_get_usdt_service/internal/otel"
	getusdt "grpc_get_usdt_service/internal/services/get_usdt"
	"grpc_get_usdt_service/internal/storage/postgresql"

	"go.uber.org/zap"
)

type App struct {
	GRPCServer *grpcapp.App

	storage *postgresql.Storage
}

func New(log *zap.Logger, cfg *config.Config, grpcPort int, otelProvider *otel.OtelProvider) *App {
	storage, err := postgresql.New(cfg)
	if err != nil {
		panic(err)
	}

	service := getusdt.New(log, storage, otelProvider)

	grpcApp := grpcapp.New(log, service, grpcPort)

	metrics.StartMetrics()

	return &App{
		GRPCServer: grpcApp,
		storage:    storage,
	}
}

func (a *App) Stop() {
	a.GRPCServer.Stop()
	a.storage.Stop()
}
