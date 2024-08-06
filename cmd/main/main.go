package main

import (
	"grpc_get_usdt_service/internal/app"
	"grpc_get_usdt_service/internal/config"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadConfig()

	log := setupLogger(cfg.Env)

	log.Info("config file: ", zap.Any("cfg", cfg))

	application := app.New(log, cfg.GRPC.Port)

	go application.GRPCServer.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()

	log.Info("application stopped")
}

func setupLogger(env string) *zap.Logger {
	switch env {
	case "local":
		if logger, err := zap.NewDevelopment(); err != nil {
			panic("error during setup logger: " + err.Error())
		} else {
			return logger
		}
	case "prod":
		if logger, err := zap.NewProduction(); err != nil {
			panic("error during setup logger: " + err.Error())
		} else {
			return logger
		}
	}

	panic("error during setup logger: env var is not set")
}
