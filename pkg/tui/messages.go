package tui

import (
	"github.com/revrost/pony/pkg/account"
	"github.com/revrost/pony/pkg/broker"
	"github.com/revrost/pony/pkg/order"
	"github.com/revrost/pony/pkg/position"
)

// Message types for the Bubble Tea update loop

type accountsLoadedMsg struct {
	accounts []*account.Account
}

type ordersLoadedMsg struct {
	orders []*order.Order
}

type positionsLoadedMsg struct {
	positions []*position.Position
}

type eventMsg struct {
	event broker.Event
}

type errMsg struct {
	err error
}
