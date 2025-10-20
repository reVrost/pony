package tui

import "github.com/revrost/pony/internal/domain"

// Message types for the Bubble Tea update loop

type accountsLoadedMsg struct {
	accounts []*domain.Account
}

type ordersLoadedMsg struct {
	orders []*domain.Order
}

type positionsLoadedMsg struct {
	positions []*domain.Position
}

type eventMsg struct {
	event domain.Event
}

type errMsg struct {
	err error
}
