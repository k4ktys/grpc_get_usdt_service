package tests

import (
	pb "grpc_get_usdt_service/protos/gen/go/get_usdt"
	"grpc_get_usdt_service/tests/suite"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRates(t *testing.T) {
	ctx, st := suite.New(t)

	res, err := st.GetUsdtClient.GetRates(ctx, &pb.GetRatesRequest{
		Market: "usdtrub",
	})

	require.NoError(t, err)

	assert.IsType(t, int64(1), res.Timestamp)
	assert.IsType(t, &pb.GetRatesResponse_Candle{}, res.Ask)
	assert.IsType(t, &pb.GetRatesResponse_Candle{}, res.Bid)

}
