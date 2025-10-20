package domain

import "time"

type Position struct {
	ID             int64
	AccountID      string
	Symbol         string
	Qty            float64
	AvgEntryPrice  float64
	CurrentPrice   float64
	MarketValue    float64
	CostBasis      float64
	UnrealizedPL   float64
	UnrealizedPLPC float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
