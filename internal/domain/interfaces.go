package domain

import "context"

// BrokerClient defines the interface for Alpaca Broker API interactions
type BrokerClient interface {
	// Account operations
	GetAccount(ctx context.Context, accountID string) (*Account, error)
	ListAccounts(ctx context.Context) ([]*Account, error)

	// Order operations
	CreateOrder(ctx context.Context, req *CreateOrderRequest) (*Order, error)
	GetOrder(ctx context.Context, orderID string) (*Order, error)
	CancelOrder(ctx context.Context, orderID string) error

	// Position operations
	ListPositions(ctx context.Context, accountID string) ([]*Position, error)

	// Event streaming
	StreamEvents(ctx context.Context, accountID string) (<-chan Event, <-chan error)
}

// Event types for SSE streaming
type EventType string

const (
	EventTypeTradeUpdate   EventType = "trade_update"
	EventTypeAccountUpdate EventType = "account_update"
)

type Event interface {
	Type() EventType
}

type TradeUpdateEvent struct {
	Order *Order
}

func (e TradeUpdateEvent) Type() EventType {
	return EventTypeTradeUpdate
}

type AccountUpdateEvent struct {
	Account *Account
}

func (e AccountUpdateEvent) Type() EventType {
	return EventTypeAccountUpdate
}
