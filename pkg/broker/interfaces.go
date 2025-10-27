package broker

import (
	"context"

	"github.com/revrost/pony/pkg/account"
	"github.com/revrost/pony/pkg/order"
	"github.com/revrost/pony/pkg/position"
)

// Client defines the interface for Alpaca Broker API interactions
type Client interface {
	// Account operations
	GetAccount(ctx context.Context, accountID string) (*account.Account, error)
	ListAccounts(ctx context.Context) ([]*account.Account, error)

	// Order operations
	CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.Order, error)
	GetOrder(ctx context.Context, orderID string) (*order.Order, error)
	CancelOrder(ctx context.Context, orderID string) error

	// Position operations
	ListPositions(ctx context.Context, accountID string) ([]*position.Position, error)

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
	Order *order.Order
}

func (e TradeUpdateEvent) Type() EventType {
	return EventTypeTradeUpdate
}

type AccountUpdateEvent struct {
	Account *account.Account
}

func (e AccountUpdateEvent) Type() EventType {
	return EventTypeAccountUpdate
}
