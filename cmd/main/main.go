package main

import (
	"grpc_get_usdt_service/internal/app"
	"grpc_get_usdt_service/internal/config"
	"grpc_get_usdt_service/internal/migrator"
	"grpc_get_usdt_service/internal/otel"
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

	migrator.Migrate("./migrations", cfg.DbUser, cfg.DbPassword, cfg.DbPort, "postgres")

	otelProvider := otel.NewOtelProvider()

	application := app.New(log, cfg, cfg.GrpcPort, otelProvider)

	go func() {
		if err := application.GRPCServer.Run(); err != nil {
			panic(err)
		}
	}()

	//gracefull shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()

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
