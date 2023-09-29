package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/quocbang/grpc-gateway/server"
	"github.com/quocbang/grpc-gateway/server/repositories/connection"
	"google.golang.org/grpc"
)

type GrpcOption struct {
	GrpcHost string
	GrpcPort int
}

func (g GrpcOption) Run() error {
	// TODO: setup interceptor here.
	s := grpc.NewServer()

	repo := connection.NewRepositories(connection.DataBaseConfig{})
	ss := server.Server{
		Repo: repo,
	}
	ss.NewServer(s)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", g.GrpcHost, g.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen in grpc, error: %v", err)
	}

	log.Printf("serving grpc host: %s, port %d \n", g.GrpcHost, g.GrpcPort)
	return s.Serve(lis)
}
