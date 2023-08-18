package logger

import (
	"context"
	"fmt"

	"golang.org/x/exp/slog"
)

var (
	Logger *slog.Logger
)

const (
	IDJsonKey string = "correlationId"
	loggerKey string = "x-logger"
)

// Init sets given logger as application default logger
func Init(logger *slog.Logger) {
	Logger = logger
}

// New adds correlationID to logger fields and embed it in given context.
func New(ctx context.Context, correlationID string) context.Context {
	fmt.Println("creating new logger with correlation id:", correlationID)
	return ToContext(ctx, Logger.With(slog.String(IDJsonKey, correlationID)))
}

// ToContext adds given logger to given context
func ToContext(ctx context.Context, logger *slog.Logger) context.Context {
	fmt.Println("adding newly created logger with loggerkey:", loggerKey)
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext returns a logger from given context. If context has no
// logger, it uses application default logger initialized in config
func FromContext(ctx context.Context) *slog.Logger {
	if ctx == nil {
		fmt.Println("returning default logger as context is nil")
		return Logger
	}

	fmt.Println("value from context", ctx.Value(loggerKey))
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		fmt.Println("got logger from context with loggerkey passed")
		return logger
	}

	fmt.Println("couldn't find logger from context with loggerkey passed")
	return Logger
}
