package main

import (
	"context"
	"fmt"
	"github.com/EgMeln/stock_market/price_service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:8081"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("can't dial price server %v", err)
	}
	client := protocol.NewPriceServiceClient(conn)

	var str []string
	str[0] = "Aeroflot"
	str[1] = "AKRON"
	req := protocol.GetRequest{Symbol: str}
	stream, err := client.Get(context.Background(), &req)
	if err != nil {
		log.Fatalf("subscrive err %v", err)
	}
	fmt.Println(stream)
}
