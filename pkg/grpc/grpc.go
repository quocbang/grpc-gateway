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
	senderConn "github.com/quocbang/grpc-gateway/server/sender/connection"
)

type GrpcOption struct {
	GrpcHost             string
	GrpcPort             int
	Database             config.DatabaseGroup
	AccessTokenLifeTime  time.Duration
	RefreshTokenLifeTime time.Duration
	SecretKey            string
	SenderConfig         config.SMTPConfig
}

func (g GrpcOption) Run() error {
	authorization := interceptors.Auth{
		SecretKey: g.SecretKey,
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			auth.UnaryServerInterceptor(authorization.Authorization),
			interceptors.UnaryLogging(),
			interceptors.UnaryRecovery(),
		),

		grpc.ChainStreamInterceptor(
			interceptors.StreamLogging(),
			interceptors.StreamRecovery(),
		),
	}
	s := grpc.NewServer(opts...)

	// init repositories layer
	connectionOpts := []connection.Options{
		connection.WithSchema(g.Database.Postgres.Schema),
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

	// init sender service.
	sender, err := senderConn.NewEmailSender(senderConn.EmailSenderConfig{
		SmtpServer:  g.SenderConfig.SmtpServer,
		SmtpPort:    g.SenderConfig.SmtpPort,
		SenderEmail: g.SenderConfig.SenderEmail,
		Password:    g.SenderConfig.Password,
	})
	if err != nil {
		return err
	}

	ss := server.Server{
		Repo:                 repo,
		Sender:               sender,
		SecretKey:            g.SecretKey,
		AccessTokenLifeTime:  g.AccessTokenLifeTime,
		RefreshTokenLifeTime: g.RefreshTokenLifeTime,
	}
	ss.NewServer(s)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", g.GrpcHost, g.GrpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen in grpc, error: %v", err)
	}

	log.Printf("serving grpc host: %s, port %d \n", g.GrpcHost, g.GrpcPort)
	return s.Serve(lis)
}
