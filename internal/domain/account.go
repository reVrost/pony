package domain

import "time"

type Account struct {
	ID              string
	AlpacaAccountID string
	Status          string
	Currency        string
	Cash            float64
	PortfolioValue  float64
	BuyingPower     float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
