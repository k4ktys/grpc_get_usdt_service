package grpcapp

import (
	"fmt"
	getusdt_grpc "grpc_get_usdt_service/internal/grpc/get_usdt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	log        *zap.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *zap.Logger, get_usdt_service getusdt_grpc.GetUsdtServicer, port int) *App {
	gRPCServer := grpc.NewServer()

	getusdt_grpc.Register(gRPCServer, get_usdt_service)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) Run() error {
	log := a.log.With(zap.String("op", "grpcapp.Run"), zap.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		return fmt.Errorf("grpcapp.Run %w", err)
	}

	log.Info("grpc server is running", zap.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("grpcapp.Run: %w", err)
	}

	return nil
}

func (a *App) Stop() {
	a.log.With(zap.String("op", "grpcapp.Stop")).Info("stopping grpc server", zap.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
