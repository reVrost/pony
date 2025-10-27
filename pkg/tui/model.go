package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/revrost/pony/pkg/account"
	"github.com/revrost/pony/pkg/broker"
	"github.com/revrost/pony/pkg/order"
	"github.com/revrost/pony/pkg/position"
)

type View int

const (
	ViewDashboard View = iota
	ViewOrders
	ViewPositions
	ViewPlaceOrder
)

// Store is the interface for database operations (will be implemented by sqlc's Querier)
type Store interface {
	// sqlc will generate these methods automatically
	// For now we just need the interface
}

type Model struct {
	currentView View
	width       int
	height      int

	// Services
	brokerClient broker.Client
	store        Store // sqlc generated Querier will implement this

	// Data
	accounts  []*account.Account
	orders    []*order.Order
	positions []*position.Position

	// State
	selectedAccount *account.Account
	err             error
	loading         bool

	// Sub-models
	placeOrderForm PlaceOrderForm
}

func NewModel(
	brokerClient broker.Client,
	store Store,
) Model {
	return Model{
		currentView:  ViewDashboard,
		brokerClient: brokerClient,
		store:        store,
		accounts:     []*account.Account{},
		orders:       []*order.Order{},
		positions:    []*position.Position{},
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		loadAccounts(m.store),
		listenForEvents(m.brokerClient),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		return m.handleKeyPress(msg)

	case accountsLoadedMsg:
		m.accounts = msg.accounts
		if len(m.accounts) > 0 {
			m.selectedAccount = m.accounts[0]
			return m, loadOrders(m.store, m.selectedAccount.ID)
		}
		return m, nil

	case ordersLoadedMsg:
		m.orders = msg.orders
		return m, nil

	case positionsLoadedMsg:
		m.positions = msg.positions
		return m, nil

	case eventMsg:
		return m.handleEvent(msg.event)

	case errMsg:
		m.err = msg.err
		m.loading = false
		return m, nil
	}

	return m, nil
}

func (m Model) View() string {
	if m.loading {
		return "Loading..."
	}

	if m.err != nil {
		return renderError(m.err)
	}

	switch m.currentView {
	case ViewDashboard:
		return renderDashboard(m)
	case ViewOrders:
		return renderOrders(m)
	case ViewPositions:
		return renderPositions(m)
	case ViewPlaceOrder:
		return renderPlaceOrder(m)
	default:
		return "Unknown view"
	}
}

func (m Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit

	case "1":
		m.currentView = ViewDashboard
		return m, nil

	case "2":
		m.currentView = ViewOrders
		if m.selectedAccount != nil {
			return m, loadOrders(m.store, m.selectedAccount.ID)
		}
		return m, nil

	case "3":
		m.currentView = ViewPositions
		if m.selectedAccount != nil {
			return m, loadPositions(m.store, m.selectedAccount.ID)
		}
		return m, nil

	case "n":
		if m.currentView == ViewOrders {
			m.currentView = ViewPlaceOrder
			m.placeOrderForm = NewPlaceOrderForm()
		}
		return m, nil

	case "esc":
		if m.currentView == ViewPlaceOrder {
			m.currentView = ViewOrders
		}
		return m, nil
	}

	// Handle sub-model key presses
	if m.currentView == ViewPlaceOrder {
		updatedForm, cmd := m.placeOrderForm.Update(msg)
		m.placeOrderForm = updatedForm
		return m, cmd
	}

	return m, nil
}

func (m Model) handleEvent(event broker.Event) (tea.Model, tea.Cmd) {
	switch e := event.(type) {
	case broker.TradeUpdateEvent:
		// Update order in local state
		for i, order := range m.orders {
			if order.AlpacaOrderID == e.Order.AlpacaOrderID {
				m.orders[i] = e.Order
				break
			}
		}
		return m, nil

	case broker.AccountUpdateEvent:
		// Update account in local state
		for i, account := range m.accounts {
			if account.AlpacaAccountID == e.Account.AlpacaAccountID {
				m.accounts[i] = e.Account
				break
			}
		}
		return m, nil
	}

	return m, nil
}
