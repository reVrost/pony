package domain

import "time"

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
	OrderStatusNew            OrderStatus = "new"
	OrderStatusPartiallyFilled OrderStatus = "partially_filled"
	OrderStatusFilled         OrderStatus = "filled"
	OrderStatusCanceled       OrderStatus = "canceled"
	OrderStatusRejected       OrderStatus = "rejected"
)

const (
	TimeInForceDay TimeInForce = "day"
	TimeInForceGTC TimeInForce = "gtc"
	TimeInForceIOC TimeInForce = "ioc"
	TimeInForceFOK TimeInForce = "fok"
)

type Order struct {
	ID              string
	AlpacaOrderID   string
	AccountID       string
	Symbol          string
	Side            OrderSide
	OrderType       OrderType
	Qty             float64
	FilledQty       float64
	LimitPrice      *float64
	StopPrice       *float64
	TimeInForce     TimeInForce
	Status          OrderStatus
	FilledAvgPrice  *float64
	SubmittedAt     *time.Time
	FilledAt        *time.Time
	CanceledAt      *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type CreateOrderRequest struct {
	AccountID   string
	Symbol      string
	Side        OrderSide
	OrderType   OrderType
	Qty         float64
	LimitPrice  *float64
	StopPrice   *float64
	TimeInForce TimeInForce
}
