// Package model contain model of struct
package model

import "github.com/google/uuid"

// Price struct that contain record info about price
type Price struct {
	ID       uuid.UUID
	Ask      float64
	Bid      float64
	Symbol   string
	DoteTime string
}
