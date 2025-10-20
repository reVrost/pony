package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/revrost/pony/internal/domain"
)

type AlpacaClient struct {
	apiKey     string
	apiSecret  string
	baseURL    string
	httpClient *http.Client
}

func NewAlpacaClient(apiKey, apiSecret, baseURL string) *AlpacaClient {
	return &AlpacaClient{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		baseURL:   baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
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
func (c *AlpacaClient) GetAccount(ctx context.Context, accountID string) (*domain.Account, error) {
	// TODO: Implement actual Alpaca Broker API call
	// GET /v1/accounts/{account_id}
	resp, err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/v1/accounts/%s", accountID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get account: status %d", resp.StatusCode)
	}

	var account domain.Account
	if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
		return nil, err
	}

	return &account, nil
}

// ListAccounts lists all accounts from Alpaca Broker API
func (c *AlpacaClient) ListAccounts(ctx context.Context) ([]*domain.Account, error) {
	// TODO: Implement actual Alpaca Broker API call
	// GET /v1/accounts
	resp, err := c.doRequest(ctx, http.MethodGet, "/v1/accounts", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list accounts: status %d", resp.StatusCode)
	}

	var accounts []*domain.Account
	if err := json.NewDecoder(resp.Body).Decode(&accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

// CreateOrder creates a new order via Alpaca Broker API
func (c *AlpacaClient) CreateOrder(ctx context.Context, req *domain.CreateOrderRequest) (*domain.Order, error) {
	// TODO: Implement actual Alpaca Broker API call
	// POST /v1/trading/accounts/{account_id}/orders
	return nil, fmt.Errorf("not implemented")
}

// GetOrder retrieves order information from Alpaca Broker API
func (c *AlpacaClient) GetOrder(ctx context.Context, orderID string) (*domain.Order, error) {
	// TODO: Implement actual Alpaca Broker API call
	// GET /v1/trading/accounts/{account_id}/orders/{order_id}
	return nil, fmt.Errorf("not implemented")
}

// CancelOrder cancels an order via Alpaca Broker API
func (c *AlpacaClient) CancelOrder(ctx context.Context, orderID string) error {
	// TODO: Implement actual Alpaca Broker API call
	// DELETE /v1/trading/accounts/{account_id}/orders/{order_id}
	return fmt.Errorf("not implemented")
}

// ListPositions lists all positions for an account from Alpaca Broker API
func (c *AlpacaClient) ListPositions(ctx context.Context, accountID string) ([]*domain.Position, error) {
	// TODO: Implement actual Alpaca Broker API call
	// GET /v1/trading/accounts/{account_id}/positions
	return nil, fmt.Errorf("not implemented")
}

// StreamEvents streams SSE events from Alpaca Broker API
func (c *AlpacaClient) StreamEvents(ctx context.Context, accountID string) (<-chan domain.Event, <-chan error) {
	eventCh := make(chan domain.Event)
	errCh := make(chan error, 1)

	go func() {
		defer close(eventCh)
		defer close(errCh)

		// TODO: Implement actual SSE streaming from Alpaca Broker API
		// This should connect to the SSE endpoint and parse events
		// Example: GET /v1/events/accounts/{account_id}

		<-ctx.Done()
		errCh <- ctx.Err()
	}()

	return eventCh, errCh
}
