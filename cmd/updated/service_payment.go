package main

import "log/slog"

type PaymentService struct {
	// stores a pointer to slog.Logger
	logger *slog.Logger
}

// "constructor" function where the logger dependency is injected.
func NewPaymentService(logger *slog.Logger) *PaymentService {
	return &PaymentService{
		logger: logger,
	}
}

func (s *PaymentService) Pay() {
	// The logger dependency isn't directly created anymore
	s.logger.Info("processing payment")
}
