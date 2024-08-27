package postgresql

import (
	"context"
	"encoding/json"
	"fmt"
	"grpc_get_usdt_service/internal/config"
	"grpc_get_usdt_service/internal/domain/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(cfg *config.Config) (*Storage, error) {
	dbConnection := fmt.Sprintf("postgres://%s:%s@get_usdt_db:%s/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbPort, cfg.DbName)

	dbpool, err := pgxpool.New(context.Background(), dbConnection)
	if err != nil {
		panic(err)
	}

	return &Storage{
		db: dbpool,
	}, nil
}

func (s *Storage) SaveUsdtTrade(ctx context.Context, trade models.UsdtTrade) error {
	data, err := json.Marshal(&trade)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, "INSERT INTO trades(trade_timestamp, data) VALUES(to_timestamp($1), $2)", trade.Timestamp, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Stop() {
	s.db.Close()
}
