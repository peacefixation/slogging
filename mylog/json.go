package mylog

import (
	"log/slog"
	"os"
)

func JSONHandler() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("Hello", "name", "Gopher", "title", "Developer")

	logger.Info("Hello", "name", "Gopher", "title")

	logger.Info("Hello again", slog.String("name", "Gopher"), slog.String("title", "Developer"))

	logger.Error("Something bad happened", "the bad thing", "ROAR")
}
