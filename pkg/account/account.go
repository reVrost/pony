package account

import (
	"time"

	"github.com/shopspring/decimal"
)

type Account struct {
	ID              string
	AlpacaAccountID string
	Status          string
	Currency        string
	Cash            decimal.Decimal
	PortfolioValue  decimal.Decimal
	BuyingPower     decimal.Decimal
	CreatedAt       time.Time
}
