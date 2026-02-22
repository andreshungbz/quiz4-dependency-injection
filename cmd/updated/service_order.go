package main

import "log/slog"

type OrderService struct {
	// stores a pointer to slog.Logger
	logger *slog.Logger
}

// "constructor" function where the logger dependency is injected.
func NewOrderService(logger *slog.Logger) *OrderService {
	return &OrderService{
		logger: logger,
	}
}

func (s *OrderService) Process() {
	// The logger dependency isn't directly created anymore
	s.logger.Info("processing order")
}
