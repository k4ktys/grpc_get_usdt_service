package getusdt

import (
	"context"
	"encoding/json"
	"grpc_get_usdt_service/internal/domain/models"
	"grpc_get_usdt_service/internal/otel"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
)

var (
	getRatesCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_usdt_service_get_rates_total",
		Help: "The total number of GetRates requests",
	})
)

type GetUsdtService struct {
	log          *zap.Logger
	tradeSaver   UsdtTradeSaver
	otelProvider *otel.OtelProvider
}

func (g *GetUsdtService) GetRates(ctx context.Context, market string) (models.UsdtTrade, error) {
	// md, _ := metadata.FromIncomingContext(ctx)
	// traceIdString := md["x-trace-id"][0]

	// traceId, err := trace.TraceIDFromHex(traceIdString)
	// if err != nil {
	// 	g.log.Error("TraceIDFromHex error", zap.String("GetRates", err.Error()))
	// } else {
	// 	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
	// 		TraceID: traceId,
	// 	})

	// 	ctx = trace.ContextWithSpanContext(ctx, spanContext)
	// }

	getRatesCount.Inc()

	ctx, span := g.otelProvider.Tracer.Start(ctx, "GetRates")
	defer span.End()

	var trades models.UsdtTrade

	res, err := http.Get("https://garantex.org/api/v2/depth?market=" + market)
	if err != nil {
		return trades, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&trades)
	if err != nil {
		return trades, err
	}

	err = g.tradeSaver.SaveUsdtTrade(ctx, trades)
	if err != nil {
		g.log.Error("error during save trade", zap.String("GetRates", err.Error()))
	}

	return trades, nil
}

type UsdtTradeSaver interface {
	SaveUsdtTrade(ctx context.Context, trade models.UsdtTrade) error
}

func New(log *zap.Logger, tradeSaver UsdtTradeSaver, otelProvider *otel.OtelProvider) *GetUsdtService {
	return &GetUsdtService{
		log:          log,
		tradeSaver:   tradeSaver,
		otelProvider: otelProvider,
	}
}
