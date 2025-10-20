package tui

import (
	"context"

	"github.com/charmbracelet/bubbletea"
	"github.com/revrost/pony/internal/domain"
)

// Commands for async operations
// These will use sqlc generated methods once we run `sqlc generate`

func loadAccounts(store Store) tea.Cmd {
	return func() tea.Msg {
		// TODO: Use store.ListAccounts() once sqlc generates it
		// For now, return empty list
		return accountsLoadedMsg{accounts: []*domain.Account{}}
	}
}

func loadOrders(store Store, accountID string) tea.Cmd {
	return func() tea.Msg {
		// TODO: Use store.ListOrders() once sqlc generates it
		// For now, return empty list
		return ordersLoadedMsg{orders: []*domain.Order{}}
	}
}

func loadPositions(store Store, accountID string) tea.Cmd {
	return func() tea.Msg {
		// TODO: Use store.ListPositions() once sqlc generates it
		// For now, return empty list
		return positionsLoadedMsg{positions: []*domain.Position{}}
	}
}

func listenForEvents(client domain.BrokerClient) tea.Cmd {
	return func() tea.Msg {
		// This is a simplified event listener
		// In a real implementation, you'd want to handle context properly
		ctx := context.Background()
		eventCh, errCh := client.StreamEvents(ctx, "")

		select {
		case event := <-eventCh:
			return eventMsg{event: event}
		case err := <-errCh:
			if err != nil {
				return errMsg{err: err}
			}
		}

		return nil
	}
}
