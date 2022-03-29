package repository

import (
	"context"
	"github.com/EgMeln/stock_market/position_service/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresPrice struct {
	PoolPrice *pgxpool.Conn
}

type PriceTransaction interface {
	OpenPosition(ctx context.Context, trans *model.Transaction) (*uuid.UUID, error)
	ClosePosition(ctx context.Context, closePrice *float64, id *uuid.UUID) error
}
