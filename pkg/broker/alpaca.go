package broker

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/revrost/pony/pkg/account"
	"github.com/revrost/pony/pkg/order"
	"github.com/revrost/pony/pkg/position"
)

type AlpacaClient struct {
	apiKey       string
	apiSecret    string
	baseURL      string
	alpacaClient *alpaca.Client
	httpClient   *http.Client
}

func NewAlpacaClient(apiKey, apiSecret, baseURL string) *AlpacaClient {
	return &AlpacaClient{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		baseURL:   baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		alpacaClient: alpaca.NewClient(alpaca.ClientOpts{
			APIKey:    apiKey,
			APISecret: apiSecret,
			BaseURL:   baseURL,
		}),
	}
}

func (c *AlpacaClient) doRequest(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("APCA-API-KEY-ID", c.apiKey)
	req.Header.Set("APCA-API-SECRET-KEY", c.apiSecret)
	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req)
}

// GetAccount retrieves account information from Alpaca Broker API
func (c *AlpacaClient) GetAccount(ctx context.Context, accountID string) (*account.Account, error) {
	resp, err := c.alpacaClient.GetAccount()
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	return &account.Account{
		ID:              resp.ID,
		AlpacaAccountID: resp.AccountNumber,
		Status:          resp.Status,
		Currency:        resp.Currency,
		Cash:            resp.Cash,
		PortfolioValue:  resp.PortfolioValue,
		BuyingPower:     resp.BuyingPower,
		CreatedAt:       resp.CreatedAt,
	}, nil
}

// ListAccounts lists all accounts from Alpaca Broker API
func (c *AlpacaClient) ListAccounts(ctx context.Context) ([]*account.Account, error) {
	// TODO: implement actual Alpaca Broker API call GetAllAccounts
	resp := make([]*alpaca.Account, 0)

	var accounts []*account.Account
	for _, acc := range resp {
		accounts = append(accounts, &account.Account{
			ID:              acc.ID,
			AlpacaAccountID: acc.AccountNumber,
			Status:          acc.Status,
			Currency:        acc.Currency,
			Cash:            acc.Cash,
			PortfolioValue:  acc.PortfolioValue,
			BuyingPower:     acc.BuyingPower,
			CreatedAt:       acc.CreatedAt,
		})
	}

	return accounts, nil
}

// CreateOrder creates a new order via Alpaca Broker API
func (c *AlpacaClient) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.Order, error) {
	resp, err := c.alpacaClient.PlaceOrder(alpaca.PlaceOrderRequest{
		Symbol:         req.Symbol,
		Qty:            req.Qty,
		Side:           alpaca.Side(req.Side),
		Type:           alpaca.OrderType(req.OrderType),
		TimeInForce:    alpaca.TimeInForce(req.TimeInForce),
		LimitPrice:     req.LimitPrice,
		ExtendedHours:  false,
		StopPrice:      req.StopPrice,
		ClientOrderID:  req.AccountID,
		OrderClass:     alpaca.OrderClass(req.OrderType),
		PositionIntent: alpaca.PositionIntent(req.OrderType),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	return OrderFromAlpaca(resp), nil
}

func OrderTypeFromAlpaca(orderType alpaca.OrderType) order.OrderType {
	switch orderType {
	case "market":
		return order.OrderTypeMarket
	case "limit":
		return order.OrderTypeLimit
	case "stop":
		return order.OrderTypeStop
	case "stop_limit":
		return order.OrderTypeStopLimit
	default:
		return order.OrderTypeMarket
	}
}

func OrderSideFromAlpaca(side alpaca.Side) order.OrderSide {
	switch side {
	case "buy":
		return order.OrderSideBuy
	case "sell":
		return order.OrderSideSell
	default:
		return order.OrderSideBuy
	}
}

func TimeInForceFromAlpaca(timeInForce alpaca.TimeInForce) order.TimeInForce {
	switch timeInForce {
	case "day":
		return order.TimeInForceDay
	case "gtc":
		return order.TimeInForceGTC
	case "ioc":
		return order.TimeInForceIOC
	case "fok":
		return order.TimeInForceFOK
	default:
		return order.TimeInForceDay
	}
}

func OrderStatusFromAlpaca(status string) order.OrderStatus {
	switch status {
	case "new":
		return order.OrderStatusNew
	case "partially_filled":
		return order.OrderStatusPartiallyFilled
	case "filled":
		return order.OrderStatusFilled
	case "canceled":
		return order.OrderStatusCanceled
	case "rejected":
		return order.OrderStatusRejected
	default:
		return order.OrderStatusNew
	}
}

func OrderFromAlpaca(o *alpaca.Order) *order.Order {
	return &order.Order{
		ID:             o.ID,
		StakeOrderID:   o.ClientOrderID,
		CreatedAt:      o.CreatedAt,
		UpdatedAt:      o.UpdatedAt,
		SubmittedAt:    o.SubmittedAt,
		FilledAt:       o.FilledAt,
		ExpiredAt:      o.ExpiredAt,
		CanceledAt:     o.CanceledAt,
		FailedAt:       o.FailedAt,
		ReplacedAt:     o.ReplacedAt,
		Symbol:         o.Symbol,
		OrderType:      OrderTypeFromAlpaca(o.Type),
		Side:           OrderSideFromAlpaca(o.Side),
		TimeInForce:    TimeInForceFromAlpaca(o.TimeInForce),
		Status:         OrderStatusFromAlpaca(o.Status),
		Qty:            o.Qty,
		FilledQty:      o.FilledQty,
		FilledAvgPrice: o.FilledAvgPrice,
		LimitPrice:     o.LimitPrice,
		StopPrice:      o.StopPrice,

		// OPTIONAL:
		// ReplacedBy:     resp.ReplacedBy,
		// Replaces:       resp.Replaces,
		// OrderClass:     resp.OrderClass,
		// AssetID:        resp.AssetID,
		// AssetClass:     resp.AssetClass,
		// PositionIntent: PositionIntentFromAlpaca(resp.PositionIntent),
		// TrailPrice:     resp.TrailPrice,
		// TrailPercent:   resp.TrailPercent,
		// HWM:            resp.HWM,
		// ExtendedHours:  resp.ExtendedHours,
		// RatioQty:       resp.RatioQty,
		// Legs:           resp.Legs,
		// Notional:       resp.Notional,
	}
}

// GetOrder retrieves order information from Alpaca Broker API
func (c *AlpacaClient) GetOrder(ctx context.Context, orderID string) (*order.Order, error) {
	resp, err := c.alpacaClient.GetOrder(orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}

	return OrderFromAlpaca(resp), nil
}

// CancelOrder cancels an order via Alpaca Broker API
func (c *AlpacaClient) CancelOrder(ctx context.Context, orderID string) error {
	return c.alpacaClient.CancelOrder(orderID)
}

// ListPositions lists all positions for an account from Alpaca Broker API
func (c *AlpacaClient) ListPositions(ctx context.Context, accountID string) ([]*position.Position, error) {
	// TODO: Implement actual Alpaca Broker API call
	// GET /v1/trading/accounts/{account_id}/positions
	return nil, fmt.Errorf("not implemented")
}

// StreamEvents streams SSE events from Alpaca Broker API
func (c *AlpacaClient) StreamEvents(ctx context.Context, accountID string) (<-chan Event, <-chan error) {
	eventCh := make(chan Event)
	errCh := make(chan error, 1)

	go func() {
		defer close(eventCh)
		defer close(errCh)

		alpaca.StreamTradeUpdatesInBackground(context.Background(), func(tu alpaca.TradeUpdate) {
			eventCh <- TradeUpdateEvent{
				Order: OrderFromAlpaca(&tu.Order),
			}
		})

		<-ctx.Done()
		errCh <- ctx.Err()
	}()

	return eventCh, errCh
}
