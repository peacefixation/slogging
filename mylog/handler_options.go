package mylog

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

// there are just ints and there are 4 levels in between each of them
// so we can insert custom levels if we want
// slog.LevelDebug -4
// slog.LevelInfo 0
// slog.LevelWarn 4
// slog.LevelError 8

const LevelTrace = slog.Level(-8) // more fine-grained than debug

// HandlerOptions shows how to use slog.HandlerOptions to configure a handler
// set the minimum level to log
// optionally add source information (file and line number)
func HandlerOptions() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug, // set the minimum level to log
		// AddSource: true, // performance penalty, internally calls runtime.Caller()
	})

	logger := slog.New(handler)

	logger.Info("some info message", "data", 123)

	logger.Debug("some low level message", "data", 456)

	logger.Log(context.Background(), LevelTrace, "a trace message", "data", 123)
}

// HandlerOptionsReplaceAttr shows how to use ReplaceAttr to customize log attributes
// customize how the level is printed
// transform the message to uppercase
func HandlerOptionsReplaceAttr() {
	opts := &slog.HandlerOptions{
		Level: LevelTrace, // set the minimum level to log
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.LevelKey:
				level := a.Value.Any().(slog.Level)
				switch level {
				case LevelTrace:
					a.Value = slog.StringValue("TRACE")
				}
			case slog.MessageKey:
				a.Value = slog.StringValue(strings.ToUpper(a.Value.String()))
			}
			return a
		},
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler)

	logger.Log(context.Background(), LevelTrace, "a trace message", "data", 123)
}
