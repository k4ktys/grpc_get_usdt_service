package getusdt

import (
	"context"
	"encoding/json"
	"grpc_get_usdt_service/internal/domain/models"
	"net/http"

	"go.uber.org/zap"
)

type GetUsdtService struct {
	log        *zap.Logger
	tradeSaver UsdtTradeSaver
}

func (g *GetUsdtService) GetRates(ctx context.Context, market string) (models.UsdtTrade, error) {
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

	err = g.tradeSaver.SaveUsdtTrade(res.Request.Context(), trades)
	if err != nil {
		g.log.Error("error during save trade", zap.String("GetRates", err.Error()))
	}

	return trades, nil
}

type UsdtTradeSaver interface {
	SaveUsdtTrade(ctx context.Context, trade models.UsdtTrade) error
}

func New(log *zap.Logger, tradeSaver UsdtTradeSaver) *GetUsdtService {
	return &GetUsdtService{
		log:        log,
		tradeSaver: tradeSaver,
	}
}
