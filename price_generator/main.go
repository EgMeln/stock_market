package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/EgMeln/stock_market/price_generator/internal/config"
	"github.com/EgMeln/stock_market/price_generator/internal/producer"
	"github.com/EgMeln/stock_market/price_generator/internal/service/generate"
	"github.com/EgMeln/stock_market/price_generator/internal/service/send"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

func main() {
	redisCfg, err := config.NewRedis()
	if err != nil {
		log.Fatalln("Config error: ", redisCfg)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB})
	if _, ok := redisClient.Ping().Result(); ok != nil {
		log.Fatalf("redis new client error %v", ok)
	}

	gen := generate.NewGenerator()
	prod := producer.NewRedis(redisClient)
	serv := send.NewService(prod, gen)
	log.Info("Start generating prices")

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	err = serv.StartSending(ctx)
	if err != nil {
		log.Fatalf("Sending error %v", err)
	}

	log.Println("received signal", <-c)
	cancel()
	err = redisClient.Close()
	if err != nil {
		log.Fatalf("redis close error %v", err)
	}
	log.Info("Success consuming messages")
}
