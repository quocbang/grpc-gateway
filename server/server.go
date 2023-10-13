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

type ServerInfo struct {
	Repo                 repositories.Repositories
	Sender               sender.Sender
	AccessTokenLifeTime  time.Duration
	RefreshTokenLifeTime time.Duration
	SecretKey            string
}

type Server struct {
	Account pb.AccountServiceServer
	Product pb.ProductServiceServer
}

func (si ServerInfo) RegisterServer() Server {
	return Server{
		Account: account.NewAccount(si.Repo, si.Sender, si.AccessTokenLifeTime, si.RefreshTokenLifeTime, si.SecretKey, roles.HasPermission),
		Product: product.NewProductService(si.Repo),
	}
}

func (sv Server) NewServer(s *grpc.Server) {
	pb.RegisterAccountServiceServer(s, sv.Account)
	pb.RegisterProductServiceServer(s, sv.Product)
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
