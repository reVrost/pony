package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("170")).
			MarginBottom(1)

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("86"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42"))

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("246"))
)

func renderError(err error) string {
	return errorStyle.Render(fmt.Sprintf("Error: %v", err))
}

func renderDashboard(m Model) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Pony Trading Terminal"))
	b.WriteString("\n\n")

	if m.selectedAccount != nil {
		b.WriteString(headerStyle.Render("Account Summary"))
		b.WriteString("\n")
		b.WriteString(fmt.Sprintf("ID: %s\n", m.selectedAccount.AlpacaAccountID))
		b.WriteString(fmt.Sprintf("Status: %s\n", m.selectedAccount.Status))
		b.WriteString(fmt.Sprintf("Cash: $%.2f\n", m.selectedAccount.Cash))
		b.WriteString(fmt.Sprintf("Portfolio Value: $%.2f\n", m.selectedAccount.PortfolioValue))
		b.WriteString(fmt.Sprintf("Buying Power: $%.2f\n", m.selectedAccount.BuyingPower))
		b.WriteString("\n")
	} else {
		b.WriteString(infoStyle.Render("No account selected"))
		b.WriteString("\n\n")
	}

	b.WriteString(renderNavigation())

	return b.String()
}

func renderOrders(m Model) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Orders"))
	b.WriteString("\n\n")

	if len(m.orders) == 0 {
		b.WriteString(infoStyle.Render("No orders found"))
		b.WriteString("\n\n")
	} else {
		b.WriteString(headerStyle.Render(fmt.Sprintf("%-15s %-10s %-6s %-10s %-12s %-15s",
			"Symbol", "Side", "Qty", "Type", "Status", "Filled")))
		b.WriteString("\n")

		for _, order := range m.orders {
			filledQty := fmt.Sprintf("%.2f/%.2f", order.FilledQty, order.Qty)
			b.WriteString(fmt.Sprintf("%-15s %-10s %-6.2f %-10s %-12s %-15s\n",
				order.Symbol,
				order.Side,
				order.Qty,
				order.OrderType,
				order.Status,
				filledQty,
			))
		}
		b.WriteString("\n")
	}

	b.WriteString(infoStyle.Render("Press 'n' to place new order"))
	b.WriteString("\n")
	b.WriteString(renderNavigation())

	return b.String()
}

func renderPositions(m Model) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Positions"))
	b.WriteString("\n\n")

	if len(m.positions) == 0 {
		b.WriteString(infoStyle.Render("No positions found"))
		b.WriteString("\n\n")
	} else {
		b.WriteString(headerStyle.Render(fmt.Sprintf("%-10s %-10s %-12s %-12s %-12s %-12s",
			"Symbol", "Qty", "Entry", "Current", "Value", "P/L")))
		b.WriteString("\n")

		for _, pos := range m.positions {
			plStyle := successStyle
			if pos.UnrealizedPL < 0 {
				plStyle = errorStyle
			}

			b.WriteString(fmt.Sprintf("%-10s %-10.2f $%-11.2f $%-11.2f $%-11.2f %s\n",
				pos.Symbol,
				pos.Qty,
				pos.AvgEntryPrice,
				pos.CurrentPrice,
				pos.MarketValue,
				plStyle.Render(fmt.Sprintf("$%.2f (%.2f%%)", pos.UnrealizedPL, pos.UnrealizedPLPC)),
			))
		}
		b.WriteString("\n")
	}

	b.WriteString(renderNavigation())

	return b.String()
}

func renderPlaceOrder(m Model) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Place Order"))
	b.WriteString("\n\n")

	b.WriteString(m.placeOrderForm.View())
	b.WriteString("\n\n")

	b.WriteString(infoStyle.Render("Press 'esc' to cancel"))
	b.WriteString("\n")

	return b.String()
}

func renderNavigation() string {
	return infoStyle.Render("\n[1] Dashboard  [2] Orders  [3] Positions  [q] Quit")
}
