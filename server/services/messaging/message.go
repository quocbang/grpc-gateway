package messaging

import (
	"context"

	"github.com/quocbang/grpc-gateway/pkg/pb"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	repo repositories.Repositories
	pb.UnimplementedMessagingServer
}

func NewMessage(repo repositories.Repositories) pb.MessagingServer {
	return server{
		repo: repo,
	}
}

func (s server) GetVerifyCode(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}
