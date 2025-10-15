package mylog

import "log/slog"

// Basic shows a simple log message using slog
// uses the default logger which logs to stdout in a text format (key=value)
func Basic() {
	slog.Info("Initializing ...", "application", "slogger demo")
}
