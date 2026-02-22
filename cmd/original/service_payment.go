package main

import (
	"log/slog"
	"os"
)

// PaymentService depends on slog.Logger. It is tightly coupled.

type PaymentService struct{}

func (s *PaymentService) Pay() {
	// dependency directly created
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("processing payment")
}
