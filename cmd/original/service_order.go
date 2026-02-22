package main

import (
	"log/slog"
	"os"
)

// OrderService depends on slog.Logger. It is tightly coupled.

type OrderService struct{}

func (s *OrderService) Process() {
	// dependency directly created
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("processing order")
}
