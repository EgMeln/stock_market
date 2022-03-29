package service

import (
	"context"
	"github.com/EgMeln/stock_market/position_service/internal/model"
	"github.com/EgMeln/stock_market/position_service/internal/repository"
	"github.com/google/uuid"
)

type PositionService struct {
	rep repository.PriceTransaction
}

func (src *PositionService) OpenPosition(ctx context.Context, trans *model.Transaction) (*uuid.UUID, error) {
	return src.rep.OpenPosition(ctx, trans)
}

func (src *PositionService) ClosePosition(ctx context.Context, closePrice *float64, id *uuid.UUID) error {
	return src.rep.ClosePosition(ctx, closePrice, id)
}
