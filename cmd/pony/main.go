package main

import (
	"database/sql"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/lib/pq"

	"github.com/revrost/pony/pkg/broker"
	"github.com/revrost/pony/pkg/config"
	"github.com/revrost/pony/pkg/tui"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Initialize database connection
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Initialize sqlc generated queries
	// After running `sqlc generate`, you'll do:
	// queries := db.New(db)
	// For now, we'll pass nil and handle it in the TUI
	var store tui.Store = nil // TODO: Replace with sqlc generated Queries

	// Initialize Alpaca broker client
	brokerClient := broker.NewAlpacaClient(
		cfg.AlpacaAPIKey,
		cfg.AlpacaAPISecret,
		cfg.AlpacaBaseURL,
	)

	// Initialize TUI model
	model := tui.NewModel(brokerClient, store)

	// Start the TUI
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running program: %w", err)
	}

	return nil
}
