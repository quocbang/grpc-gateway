package server

import (
	"context"
	"fmt"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quocbang/grpc-gateway/pkg/pb"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/sender"
	"github.com/quocbang/grpc-gateway/server/services/account"
	"github.com/quocbang/grpc-gateway/server/services/product"
	"github.com/quocbang/grpc-gateway/server/utils/roles"
	"google.golang.org/grpc"
)

type Server struct {
	Repo                 repositories.Repositories
	Sender               sender.Sender
	AccessTokenLifeTime  time.Duration
	RefreshTokenLifeTime time.Duration
	SecretKey            string
}

func (sv Server) NewServer(s *grpc.Server) {
	pb.RegisterAccountServiceServer(s, account.NewAccount(sv.Repo, sv.Sender, sv.AccessTokenLifeTime, sv.RefreshTokenLifeTime, sv.SecretKey, roles.HasPermission))
	pb.RegisterProductServiceServer(s, product.NewProductService(sv.Repo))
}

func NewRegisterHandler(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opt []grpc.DialOption) error {
	if err := pb.RegisterAccountServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opt); err != nil {
		return fmt.Errorf("failed to register account handler from end point, error: %v", err)
	}

	if err := pb.RegisterProductServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opt); err != nil {
		return fmt.Errorf("failed to register product handler from end point, error: %v", err)
	}

	return nil
}
