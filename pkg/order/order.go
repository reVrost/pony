package order

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderSide string
type OrderType string
type OrderStatus string
type TimeInForce string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

const (
	OrderTypeMarket    OrderType = "market"
	OrderTypeLimit     OrderType = "limit"
	OrderTypeStop      OrderType = "stop"
	OrderTypeStopLimit OrderType = "stop_limit"
)

const (
	OrderStatusNew             OrderStatus = "new"
	OrderStatusPartiallyFilled OrderStatus = "partially_filled"
	OrderStatusFilled          OrderStatus = "filled"
	OrderStatusCanceled        OrderStatus = "canceled"
	OrderStatusRejected        OrderStatus = "rejected"
)

const (
	TimeInForceDay TimeInForce = "day"
	TimeInForceGTC TimeInForce = "gtc"
	TimeInForceIOC TimeInForce = "ioc"
	TimeInForceFOK TimeInForce = "fok"
)

type Order struct {
	ID             string
	StakeOrderID   string
	AlpacaOrderID  string
	AccountID      string
	Symbol         string
	Side           OrderSide
	OrderType      OrderType
	Qty            *decimal.Decimal
	FilledQty      decimal.Decimal
	LimitPrice     *decimal.Decimal
	StopPrice      *decimal.Decimal
	TimeInForce    TimeInForce
	Status         OrderStatus
	FilledAvgPrice *decimal.Decimal
	SubmittedAt    time.Time
	FilledAt       *time.Time
	ExpiredAt      *time.Time
	CanceledAt     *time.Time
	FailedAt       *time.Time
	ReplacedAt     *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreateOrderRequest struct {
	AccountID   string
	Symbol      string
	Side        OrderSide
	OrderType   OrderType
	Qty         *decimal.Decimal
	LimitPrice  *decimal.Decimal
	StopPrice   *decimal.Decimal
	TimeInForce TimeInForce
}
