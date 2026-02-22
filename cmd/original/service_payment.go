package main

import (
	"log/slog"
	"os"
)

// Payment depends on slog.Logger. It is tightly coupled.

type Payment struct{}

func (s *Payment) Pay() {
	// dependency directly created
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("processing payment")
}
