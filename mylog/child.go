package mylog

import (
	"log/slog"
	"os"
)

// ChildLogger shows how to create child loggers with additional context
// use With to add context that will be included in all log entries from the child logger
// this is useful for adding component names, request IDs, etc
func ChildLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("starting up", "version", "1.0.0")

	dbLogger := logger.With("component", "database")

	dbLogger.Info("connecting to DB", "host", "localhost", "port", 5432)

	apiLogger := logger.With(slog.String("component", "api"))

	apiLogger.Info("starting API server", "address", ":8080")

	requestLogger := apiLogger.With("request_id", "abc123")

	requestLogger.Info("handling request", "method", "GET", "path", "/users")
}
