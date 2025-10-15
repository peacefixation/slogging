package mylog

import (
	"log/slog"
	"os"
)

// JSONHandler shows how to use the built-in JSON handler
// it produces structured JSON logs
// we can make mistakes by not passing key-value pairs correctly
// but there are ways to mitigate that:
// - use slog attributes
// - use a linter that checks for correct usage
func JSONHandler() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("Hello", "name", "Gopher", "title", "Developer")

	logger.Info("Hello", "name", "Gopher", "title")

	logger.Info("Hello again", slog.String("name", "Gopher"), slog.String("title", "Developer"))

	logger.Error("Something bad happened", "the bad thing", "ROAR")
}
