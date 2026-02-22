package main

import (
	"log/slog"
	"os"
)

// Order depends on slog.Logger. It is tightly coupled.

type Order struct{}

func (s *Order) Process() {
	// dependency directly created
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("processing order")
}
