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
	cfg, err := config.NewConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	log := setupLogger(cfg.Env)

	application := app.New(log, cfg, cfg.GrpcPort)

	go application.GRPCServer.Run()

	//gracefull shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()

	log.Info("application stopped")
}

func setupLogger(env string) *zap.Logger {
	switch env {
	case "dev":
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
