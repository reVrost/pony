# Pony Trading Terminal

A TUI (Terminal User Interface) stock trading application built with Go, integrated with Alpaca Broker API.

## Architecture

This project follows clean architecture principles with minimal abstractions:

```
├── cmd/pony/              # Application entry point
├── internal/
│   ├── domain/            # Domain models and interfaces (business logic)
│   ├── broker/            # Alpaca Broker API client implementation
│   ├── db/                # sqlc generated code (after running `make sqlc`)
│   └── tui/               # Bubble Tea TUI implementation
├── pkg/
│   ├── config/            # Configuration management
│   └── logger/            # Logging utilities (future)
├── db/
│   ├── schema.sql         # Database schema
│   └── queries/           # SQL queries for sqlc
└── docker-compose.yml     # Local PostgreSQL setup
```

**Why no repository layer?** sqlc generates type-safe Go code that IS the repository layer. No need for extra wrappers!

## Features

- **Clean & Simple**: No unnecessary abstractions - sqlc handles data access
- **TUI Interface**: Built with Bubble Tea for a beautiful terminal UI
- **Alpaca Broker API Integration**: Ready to integrate with Alpaca Broker API
- **Event Streaming**: SSE event listener for real-time order/account updates
- **PostgreSQL Database**: Local storage for accounts, orders, and positions
- **sqlc**: Type-safe SQL query generation - no ORM bloat

## Prerequisites

- Go 1.25.2 or later
- Docker and Docker Compose
- sqlc (install via `make install-tools`)

## Getting Started

1. **Clone and setup environment**:
   ```bash
   cp .env.example .env
   # Edit .env with your Alpaca API credentials
   ```

2. **Install tools**:
   ```bash
   make install-tools
   ```

3. **Start the development environment**:
   ```bash
   make dev
   ```
   This will:
   - Start PostgreSQL in Docker
   - Run database migrations
   - Set up the schema

4. **Generate sqlc code**:
   ```bash
   make sqlc
   ```
   This generates type-safe Go code in `internal/db/`

5. **Run the application**:
   ```bash
   make run
   ```

## Development Workflow

### Database Operations

- Start database: `make db-up`
- Stop database: `make db-down`
- Run migrations: `make db-migrate`
- Generate sqlc code: `make sqlc`

### Building

- Build binary: `make build`
- Run application: `make run`
- Clean build artifacts: `make clean`

## TUI Navigation

- `1` - Dashboard view (account summary)
- `2` - Orders view
- `3` - Positions view
- `n` - Place new order (when in Orders view)
- `esc` - Cancel/go back
- `q` or `Ctrl+C` - Quit application

## How It Works

### Data Flow

```
TUI (Bubble Tea)
    ↓
    ├─→ BrokerClient (Alpaca API) → external API calls
    └─→ Store (sqlc Querier) → database queries
```

### Why sqlc Instead of a Repository Layer?

**Before (overcomplicated):**
```go
// Repository interface
type OrderRepository interface {
    Create(ctx, order) error
}

// Repository implementation
func (r *Repo) Create(ctx, order) error {
    return r.queries.CreateOrder(...)  // Just wrapping sqlc!
}
```

**After (clean):**
```go
// Just use sqlc directly
queries := db.New(dbConn)
order, err := queries.CreateOrder(ctx, db.CreateOrderParams{...})
```

sqlc generates:
- Type-safe functions
- Proper interfaces (Querier)
- All CRUD operations
- No reflection, no magic

**You only need a wrapper layer when:**
- Doing complex cross-table transactions
- Adding caching
- Translating between DB models and domain models
- Adding audit logging

For this app? You don't need any of that. Keep it simple.

## Next Steps

This is a skeleton implementation. Here's what needs to be implemented:

### High Priority
1. **Complete Alpaca Broker API Integration**:
   - Implement all API endpoints in `internal/broker/alpaca.go`
   - Add proper request/response models
   - Implement SSE event streaming

2. **Wire up sqlc**:
   - Run `make sqlc` to generate code
   - Update `cmd/pony/main.go` to use generated queries
   - Update `internal/tui/commands.go` to call sqlc methods

3. **Enhance TUI**:
   - Add proper text input fields (use Bubble Tea components)
   - Implement order submission
   - Add error handling and notifications
   - Add loading states

### Medium Priority
4. **Event Processing**:
   - Properly handle SSE events from Alpaca
   - Update database on events
   - Update TUI in real-time

5. **Testing**:
   - Add unit tests for domain logic
   - Add integration tests
   - Add mock broker client for testing

### Low Priority
6. **Additional Features**:
   - Account switching
   - Order history filtering
   - Position P/L tracking
   - Configuration file support
   - Logging

## Project Structure Rationale

- **`internal/domain/`**: Core business entities and interfaces. No dependencies.
- **`internal/broker/`**: Alpaca API client. Implements `domain.BrokerClient`.
- **`internal/db/`**: sqlc generated code. Type-safe database operations.
- **`internal/tui/`**: UI layer. Completely separate from business logic.
- **`cmd/pony/`**: Wires everything together. Dependency injection here.

This separation allows you to:
- Test business logic independently
- Swap implementations easily (e.g., mock broker)
- Keep the TUI and backend completely decoupled
- Use sqlc's generated code directly without wrappers

## Contributing

This is a skeleton project. Feel free to:
1. Implement the TODOs in the code
2. Add new features
3. Improve error handling
4. Add tests
5. Enhance the TUI

## License

MIT
