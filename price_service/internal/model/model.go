package model

import (
	"encoding/json"
	"github.com/google/uuid"
)

type GeneratedPrice struct {
	ID       uuid.UUID
	Ask      float64
	Bid      float64
	Symbol   string
	DoteTime string
}

func (gen *GeneratedPrice) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &gen)
}
