package server

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quocbang/grpc-gateway/pkg/pb"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/services/account"
	"github.com/quocbang/grpc-gateway/server/services/messaging"
	"google.golang.org/grpc"
)

type Server struct {
	Repo repositories.Repositories
}

func (sv Server) NewServer(s *grpc.Server) {
	pb.RegisterAccountServer(s, account.NewAccount(sv.Repo))
	pb.RegisterMessagingServer(s, messaging.NewMessage(sv.Repo))
}

func NewRegisterHandler(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opt []grpc.DialOption) error {
	if err := pb.RegisterAccountHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opt); err != nil {
		return fmt.Errorf("failed to register account handler from end point, error: %v", err)
	}

	if err := pb.RegisterMessagingHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opt); err != nil {
		return fmt.Errorf("failed to register messaging handler from end point, error: %v", err)
	}

	return nil
}
