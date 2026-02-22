package main

// Problem Scenario: In this application we have two services. These
// services depend on a slog.Logger. From here, it is unclear that they
// depend on a logger. Moreover, if we wanted to change the configuration
// of the logger, we will have to do it for both services. The service and
// logger are tightly coupled.

// Example requirement changes:
// 1) Using slog.NewJSONHandler instead of slog.NewTextHandler
// 2) Adding a global attribute .With("environment", "production")

func main() {
	order := Order{}
	payment := Payment{}

	order.Process()
	payment.Pay()
}
