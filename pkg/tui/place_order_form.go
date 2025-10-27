package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbletea"
)

type PlaceOrderForm struct {
	symbol      string
	side        string
	qty         string
	orderType   string
	limitPrice  string
	stopPrice   string
	timeInForce string
	focusIndex  int
}

func NewPlaceOrderForm() PlaceOrderForm {
	return PlaceOrderForm{
		side:        "buy",
		orderType:   "market",
		timeInForce: "day",
		focusIndex:  0,
	}
}

func (f PlaceOrderForm) Update(msg tea.KeyMsg) (PlaceOrderForm, tea.Cmd) {
	switch msg.String() {
	case "tab", "down":
		f.focusIndex++
		if f.focusIndex > 6 {
			f.focusIndex = 0
		}
		return f, nil

	case "shift+tab", "up":
		f.focusIndex--
		if f.focusIndex < 0 {
			f.focusIndex = 6
		}
		return f, nil

	case "enter":
		// TODO: Submit order
		return f, nil

	default:
		// Handle text input for focused field
		return f.handleInput(msg.String()), nil
	}
}

func (f PlaceOrderForm) handleInput(input string) PlaceOrderForm {
	// Simplified input handling - in production you'd want proper text input
	switch f.focusIndex {
	case 0:
		if input == "backspace" && len(f.symbol) > 0 {
			f.symbol = f.symbol[:len(f.symbol)-1]
		} else if len(input) == 1 {
			f.symbol += input
		}
	case 2:
		if input == "backspace" && len(f.qty) > 0 {
			f.qty = f.qty[:len(f.qty)-1]
		} else if len(input) == 1 {
			f.qty += input
		}
	}
	return f
}

func (f PlaceOrderForm) View() string {
	cursor := func(active bool) string {
		if active {
			return ">"
		}
		return " "
	}

	return fmt.Sprintf(`
%s Symbol:       %s
%s Side:         %s  (TODO: toggle)
%s Quantity:     %s
%s Type:         %s  (TODO: select)
%s Limit Price:  %s
%s Stop Price:   %s
%s Time in Force: %s  (TODO: select)

Press [Enter] to submit
`,
		cursor(f.focusIndex == 0), f.symbol,
		cursor(f.focusIndex == 1), f.side,
		cursor(f.focusIndex == 2), f.qty,
		cursor(f.focusIndex == 3), f.orderType,
		cursor(f.focusIndex == 4), f.limitPrice,
		cursor(f.focusIndex == 5), f.stopPrice,
		cursor(f.focusIndex == 6), f.timeInForce,
	)
}
