package main

import "log/slog"

type Payment struct {
	// stores a pointer to slog.Logger
	logger *slog.Logger
}

// "constructor" function where the logger dependency is injected.
func NewPayment(logger *slog.Logger) *Payment {
	return &Payment{
		logger: logger,
	}
}

func (s *Payment) Pay() {
	// The logger dependency isn't directly created anymore
	s.logger.Info("processing payment")
}
