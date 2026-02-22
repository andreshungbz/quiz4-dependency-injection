package main

import (
	"log/slog"
	"os"
)

// Solution Scenario: With dependency injection implemented, our main program
// essentially wires the dependencies with those objects who need them. It is
// clearer to see that the services require depends on a logger with the
// "constructor" functions. We create the logger here and "inject" it into the
// service. The service and the logger are now loosely coupled.

// There are some immediate benefits we can see:
// 1) Reusability - in using a single logger for multiple services.
// 2) Easy Runtime Configuration - modifying the logger from main.
// 3) Flexibility - in configuring different loggers for different services.

func main() {
	// textLogger := slog.New(
	// 	slog.NewJSONHandler(os.Stdout, nil),
	// ).With("environment", "production")

	jsonLogger := slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	).With("environment", "production")

	// using the "constructor" functions to inject the dependency into the service
	orderService := NewOrderService(jsonLogger)
	paymentService := NewPaymentService(jsonLogger)

	orderService.Process()
	paymentService.Pay()
}
