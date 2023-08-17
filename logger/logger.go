package logger

import (
	"context"
	"log/slog"
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
	return ToContext(ctx, Logger.With(slog.String(IDJsonKey, correlationID)))
}

// ToContext adds given logger to given context
func ToContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext returns a logger from given context. If context has no
// logger, it uses application default logger initialized in config
func FromContext(ctx context.Context) *slog.Logger {
	if ctx == nil {
		return Logger
	}

	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return logger
	}

	return Logger
}
