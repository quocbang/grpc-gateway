package messaging

import (
	"context"

	"github.com/quocbang/grpc-gateway/pkg/pb"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	repo repositories.Repositories
	pb.UnimplementedMessagingServiceServer
}

func NewMessage(repo repositories.Repositories) pb.MessagingServiceServer {
	return server{
		repo: repo,
	}
}

func (s server) GetVerifyCode(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}
