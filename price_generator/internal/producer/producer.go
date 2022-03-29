// Package producer contain function for produce messages to redis
package producer

import (
	"fmt"

	"github.com/EgMeln/stock_market/price_generator/internal/service/generate"
	"github.com/go-redis/redis"
)

// Producer struct for redis client
type Producer struct {
	redisClient *redis.Client
	redisStream string
}

// NewRedis returns new instance of Producer
func NewRedis(cln *redis.Client) *Producer {
	return &Producer{redisClient: cln, redisStream: "STREAM"}
}

// ProduceMessage send messages to redis stream
func (cln *Producer) ProduceMessage(prices *generate.Generator) error {
	err := cln.redisClient.XAdd(&redis.XAddArgs{
		Stream: cln.redisStream,
		Values: prices.Prices,
	}).Err()
	if err != nil {
		return fmt.Errorf("redis send message error %w", err)
	}
	return nil
}
