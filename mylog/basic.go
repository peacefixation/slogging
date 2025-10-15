package mylog

import "log/slog"

func Basic() {
	// uses default logger which logs to stdout in a text format (key=value)
	slog.Info("Initializing ...", "application", "slogger demo")
}
