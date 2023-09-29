package account

import (
	"context"

	"github.com/quocbang/grpc-gateway/pkg/pb"
	"github.com/quocbang/grpc-gateway/server/repositories"
)

type server struct {
	repo repositories.Repositories
	pb.UnimplementedAccountServer
}

func NewAccount(repo repositories.Repositories) pb.AccountServer {
	return server{
		repo: repo,
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
