// Package send contain function for process
package send

import (
	"context"
	"time"

	"github.com/EgMeln/stock_market/price_generator/internal/producer"
	"github.com/EgMeln/stock_market/price_generator/internal/service/generate"
	log "github.com/sirupsen/logrus"
)

// Service struct
type Service struct {
	prod   *producer.Producer
	prices *generate.Generator
}

// NewService create new service
func NewService(prod *producer.Producer, prices *generate.Generator) *Service {
	return &Service{prod: prod, prices: prices}
}

// StartSending start generate prices
func (serv *Service) StartSending(ctx context.Context) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				serv.prices.GeneratePrices()
				err := serv.prod.ProduceMessage(serv.prices)
				if err != nil {
					log.Errorf("Error with sending message %v", err)
					return
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return nil
}
