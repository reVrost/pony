# Architecture Overview

### Clean Separation

```
┌─────────────────────────────────────┐
│           TUI (Bubble Tea)           │
│  - Views (rendering)                 │
│  - Model (state)                     │
│  - Commands (async ops)              │
└──────────┬──────────────┬───────────┘
           │              │
           ▼              ▼
    ┌──────────┐    ┌──────────┐
    │  Broker  │    │  Store   │
    │ (Alpaca) │    │  (sqlc)  │
    └──────────┘    └──────────┘
         │                │
         ▼                ▼
    External API      Database
```

### Domain Layer

Contains:

- **Models**: `Account`, `Order`, `Position`
- **Interfaces**: `broker.Client`
- **Events**: SSE event types

No database interfaces needed - sqlc handles that.
