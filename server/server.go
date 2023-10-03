package server

import (
	"context"
	"fmt"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quocbang/grpc-gateway/pkg/pb"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/services/account"
	"github.com/quocbang/grpc-gateway/server/services/messaging"
	"github.com/quocbang/grpc-gateway/server/services/product"
	"github.com/quocbang/grpc-gateway/server/utils/roles"
	"google.golang.org/grpc"
)

type Server struct {
	Repo          repositories.Repositories
	TokenLifeTime time.Duration
	SecretKey     string
}

func (sv Server) NewServer(s *grpc.Server) {
	pb.RegisterAccountServiceServer(s, account.NewAccount(sv.Repo, sv.TokenLifeTime, sv.SecretKey, roles.HasPermission))
	pb.RegisterMessagingServiceServer(s, messaging.NewMessage(sv.Repo))
	pb.RegisterProductServiceServer(s, product.NewProductService(sv.Repo))
}

func NewRegisterHandler(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opt []grpc.DialOption) error {
	if err := pb.RegisterAccountServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opt); err != nil {
		return fmt.Errorf("failed to register account handler from end point, error: %v", err)
	}

	if err := pb.RegisterMessagingServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opt); err != nil {
		return fmt.Errorf("failed to register messaging handler from end point, error: %v", err)
	}

	return nil
}
