package grpc

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"

	"github.com/quocbang/grpc-gateway/config"
	"github.com/quocbang/grpc-gateway/pkg/grpc/interceptors"
	"github.com/quocbang/grpc-gateway/server"
	"github.com/quocbang/grpc-gateway/server/repositories/connection"
)

type GrpcOption struct {
	GrpcHost      string
	GrpcPort      int
	Database      config.DatabaseGroup
	TokenLifeTime time.Duration
	SecretKey     string
}

func (g GrpcOption) Run() error {
	// TODO: setup interceptor here.
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			auth.UnaryServerInterceptor(interceptors.Authorization),
			interceptors.UnaryLogging(),
		),

		grpc.ChainStreamInterceptor(
			interceptors.StreamLogging(),
		),
	}
	s := grpc.NewServer(opts...)

	// init repositories layer
	connectionOpts := []connection.Options{
		connection.WithScheme(g.Database.Postgres.Schema),
		connection.WithAutoMigrate(),
	}
	repo, err := connection.NewRepositories(connection.DataBaseConfig{
		Postgres: connection.PostgresConfig{
			Address:  g.Database.Postgres.Address,
			Port:     g.Database.Postgres.Port,
			Name:     g.Database.Postgres.Name,
			Username: g.Database.Postgres.UserName,
			Password: g.Database.Postgres.Password,
		},
	}, connectionOpts...)
	if err != nil {
		return err
	}

	ss := server.Server{
		Repo:          repo,
		TokenLifeTime: g.TokenLifeTime,
	}
	ss.NewServer(s)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", g.GrpcHost, g.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen in grpc, error: %v", err)
	}

	log.Printf("serving grpc host: %s, port %d \n", g.GrpcHost, g.GrpcPort)
	return s.Serve(lis)
}
