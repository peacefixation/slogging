package mylog

import (
	"log/slog"
	"os"
)

func ChildLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("starting up", "version", "1.0.0")

	dbLogger := logger.With("component", "database")

	dbLogger.Info("connecting to DB", "host", "localhost", "port", 5432)

	apiLogger := logger.With(slog.String("component", "api"))

	apiLogger.Info("starting API server", "address", ":8080")
}
