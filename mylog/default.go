package mylog

import (
	"log/slog"
	"os"
)

// DefaultLogger shows how to set the default logger
func DefaultLogger() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	slog.Info("starting up", "version", "1.0.0")
}
