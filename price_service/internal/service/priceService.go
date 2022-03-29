package service

import (
	"github.com/EgMeln/stock_market/price_service/internal/model"
	"github.com/EgMeln/stock_market/price_service/protocol"
	"sync"
)

type PriceServer struct {
	mu           *sync.RWMutex
	generatedMap *map[string]*model.GeneratedPrice
	protocol.UnimplementedPriceServiceServer
}

func NewPriceServer(mu *sync.RWMutex, priceMap *map[string]*model.GeneratedPrice) *PriceServer {
	return &PriceServer{generatedMap: priceMap, mu: mu}
}

func (priceServ *PriceServer) Send(in *protocol.GetRequest, stream protocol.PriceService_GetServer) error {
	key := in.Symbol[0]
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			priceServ.mu.Lock()
			resp := (*priceServ.generatedMap)[key]
			priceServ.mu.Unlock()
			cur := protocol.Price{Symbol: resp.Symbol, Ask: float32(resp.Ask), Bid: float32(resp.Bid), ID: resp.ID.String(), Time: resp.DoteTime}
			err := stream.Send(&protocol.GetResponse{Price: &cur})
			if err != nil {
				return err
			}
		}
	}
}
