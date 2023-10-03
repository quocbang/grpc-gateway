package interceptors

import (
	"context"

	"go.uber.org/zap"
)

type key int

const (
	loggerKey key = iota + 1
)

func setContextWithLogger(parent context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(parent, loggerKey, logger)
}

func GetLoggerFormContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return logger
	}

	return zap.L()
}
