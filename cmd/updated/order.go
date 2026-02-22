package main

import "log/slog"

type Order struct {
	// stores a pointer to slog.Logger
	logger *slog.Logger
}

// "constructor" function where the logger dependency is injected.
func NewOrder(logger *slog.Logger) *Order {
	return &Order{
		logger: logger,
	}
}

func (s *Order) Process() {
	// The logger dependency isn't directly created anymore
	s.logger.Info("processing order")
}
