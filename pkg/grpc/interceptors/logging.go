package interceptors

import (
	"context"
	"strings"
	"time"

	"github.com/rs/xid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	mtd "github.com/quocbang/grpc-gateway/pkg/metadata"
)

// getRequestID is get request id from other service if exist, if not generate new.
func getRequestID(md metadata.MD) string {
	requestIDs := md.Get(mtd.RequestID)
	if len(requestIDs) > 0 {
		return strings.ToLower(requestIDs[0])
	}

	return xid.New().String()
}

func UnaryLogging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		md, _ := metadata.FromIncomingContext(ctx)
		logger := zap.L().With(
			zap.String("path", info.FullMethod),
			zap.String("request_id", getRequestID(md)),
		)
		newCtx := setContextWithLogger(ctx, logger)

		now := time.Now()
		logger.Info("start unary call")
		resp, err = handler(newCtx, req)
		if err != nil {
			logger.Error("finished unary call with error",
				zap.Any("request", req),
				zap.Any("response", resp),
				zap.Error(err),
				zap.Duration("elapsed_time", time.Since(now)),
			)
			return resp, err
		}

		logger.Info("finished unary call",
			zap.Any("response", resp),
			zap.Duration("elapsed_time", time.Since(now)),
		)

		return resp, err
	}
}

func StreamLogging() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		logger := zap.L().With(
			zap.String("path", info.FullMethod),
		)

		// stream.ctx = setContextWithLogger(ctx, logger)
		logger.Info("start stream call")

		now := time.Now()
		err := handler(srv, ss)
		if err != nil {
			logger.Error("finished stream call with error",
				zap.Any("request", srv),
				zap.Duration("elapsed_time", time.Since(now)),
				zap.Error(err),
			)
			return err
		}
		logger.Info("finished stream call",
			zap.Duration("elapsed_time", time.Since(now)))
		return err
	}
}
