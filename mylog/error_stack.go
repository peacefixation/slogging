package mylog

import (
	"context"
	"log/slog"
	"os"

	xerrors "github.com/mdobak/go-xerrors"
)

func ErrorStackLogger() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: replaceAttr,
	})

	logger := slog.New(handler)

	err := doSomeWork()

	logger.ErrorContext(context.Background(), "doing some work", slog.Any("error", err))
}

func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	if err, ok := a.Value.Any().(error); ok {
		if trace := xerrors.StackTrace(err); len(trace) > 0 {
			errGroup := slog.GroupValue(
				slog.String("msg", err.Error()),
				slog.Any("trace", formatStackTrace(trace)),
			)
			a.Value = errGroup
		}
	}
	return a
}

func formatStackTrace(trace xerrors.Callers) []map[string]any {
	frames := trace.Frames()
	s := make([]map[string]any, len(frames))
	for i, v := range frames {
		s[i] = map[string]any{
			"func":   v.Function,
			"source": v.File,
			"line":   v.Line,
		}
	}
	return s
}

func doSomeWork() error {
	return xerrors.New("something bad happened")
}
