package account

import (
	"context"
	"time"

	"github.com/quocbang/grpc-gateway/pkg/pb"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/utils/roles"
)

type server struct {
	repo          repositories.Repositories
	tokenLifeTime time.Duration
	secretKey     string
	hasPermission func(string, roles.Roles) bool
	pb.UnimplementedAccountServiceServer
}

func NewAccount(repo repositories.Repositories, tokenLifeTime time.Duration, secretKey string, hasPermission func(string, roles.Roles) bool) pb.AccountServiceServer {
	return server{
		repo:          repo,
		tokenLifeTime: tokenLifeTime,
		secretKey:     secretKey,
		hasPermission: hasPermission,
	}
}

func (s server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		AccessToken:  req.Username,
		RefreshToken: req.Password,
	}, nil
}

func (s server) ReNewToken(ctx context.Context, req *pb.ReNewTokenRequest) (*pb.ReNewTokenResponse, error) {
	return nil, nil
}
