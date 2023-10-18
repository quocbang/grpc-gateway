package interceptors

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

	"github.com/quocbang/grpc-gateway/server/utils/token"
)

type key int

const (
	loggerKey key = iota + 1
	JWTAuth
)

const (
	UserAgent = "GRPCGATEWAY-User-Agent"
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

func GetUserAgentFromContext(ctx context.Context) (string, error) {
	MD, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("failed to get metadata")
	}

	userAgents := MD.Get(UserAgent)
	if len(userAgents) > 0 {
		return userAgents[0], nil
	}

	return "", fmt.Errorf("failed to get user agent")
}

func GetClientIPFromContext(ctx context.Context) (string, error) {
	peer, _ := peer.FromContext(ctx)

	host, port, err := net.SplitHostPort(peer.Addr.String())
	if err != nil {
		return "", fmt.Errorf("failed to split host and port, error: %v", err)
	}

	return fmt.Sprintf("%s:%s", host, port), nil
}
