-- Schema for Pony Trading App

CREATE TABLE IF NOT EXISTS accounts (
    id TEXT PRIMARY KEY,
    alpaca_account_id TEXT UNIQUE NOT NULL,
    status TEXT NOT NULL,
    currency TEXT NOT NULL,
    cash DECIMAL(20, 2) NOT NULL DEFAULT 0,
    portfolio_value DECIMAL(20, 2) NOT NULL DEFAULT 0,
    buying_power DECIMAL(20, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS orders (
    id TEXT PRIMARY KEY,
    alpaca_order_id TEXT UNIQUE NOT NULL,
    account_id TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    symbol TEXT NOT NULL,
    side TEXT NOT NULL, -- buy or sell
    order_type TEXT NOT NULL, -- market, limit, stop, stop_limit
    qty DECIMAL(20, 8) NOT NULL,
    filled_qty DECIMAL(20, 8) NOT NULL DEFAULT 0,
    limit_price DECIMAL(20, 2),
    stop_price DECIMAL(20, 2),
    time_in_force TEXT NOT NULL, -- day, gtc, ioc, fok
    status TEXT NOT NULL, -- new, partially_filled, filled, canceled, rejected
    filled_avg_price DECIMAL(20, 2),
    submitted_at TIMESTAMP,
    filled_at TIMESTAMP,
    canceled_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_orders_account_id ON orders(account_id);
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_orders_symbol ON orders(symbol);
CREATE INDEX idx_orders_created_at ON orders(created_at DESC);

CREATE TABLE IF NOT EXISTS positions (
    id SERIAL PRIMARY KEY,
    account_id TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    symbol TEXT NOT NULL,
    qty DECIMAL(20, 8) NOT NULL,
    avg_entry_price DECIMAL(20, 2) NOT NULL,
    current_price DECIMAL(20, 2) NOT NULL,
    market_value DECIMAL(20, 2) NOT NULL,
    cost_basis DECIMAL(20, 2) NOT NULL,
    unrealized_pl DECIMAL(20, 2) NOT NULL,
    unrealized_plpc DECIMAL(10, 4) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(account_id, symbol)
);

CREATE INDEX idx_positions_account_id ON positions(account_id);
