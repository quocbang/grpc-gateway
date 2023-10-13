package interceptors

import (
	"context"

	"go.uber.org/zap"

	"github.com/quocbang/grpc-gateway/server/utils/token"
)

type key int

const (
	loggerKey key = iota + 1
	JWTAuth
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

func setContextWithJWTClaims(parent context.Context, claims *token.JWTClaimCustom) context.Context {
	return context.WithValue(parent, JWTAuth, claims)
}

func GetJWTClaimsFromContext(ctx context.Context) *token.JWTClaimCustom {
	if claims, ok := ctx.Value(JWTAuth).(*token.JWTClaimCustom); ok {
		return claims
	}
	return nil
}
