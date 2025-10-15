package mylog

import (
	"errors"
	"log/slog"
	"os"
)

type CustomError struct {
	Code int
	Msg  string
}

func (e CustomError) Error() string {
	return e.Msg
}

type CustomErrorLogValuer struct {
	Code int
	Msg  string
}

func (e CustomErrorLogValuer) Error() string {
	return e.Msg
}

func (e CustomErrorLogValuer) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("code", e.Code),
		slog.String("message", e.Msg),
	)
}

// ErrorLogger logs errors, including custom error types that implement the LogValue method
func ErrorLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	var err = errors.New("a basic error")

	logger.Error("basic error", "error", err)

	logger.Error("custom error", "error", CustomError{Code: 500, Msg: "database connection failed"})

	logger.Error("custom error log valuer", "error", CustomErrorLogValuer{Code: 500, Msg: "database connection failed"})
}
