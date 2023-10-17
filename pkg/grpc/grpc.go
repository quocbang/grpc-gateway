package grpc

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/hibiken/asynq"
	"google.golang.org/grpc"

	"github.com/quocbang/grpc-gateway/config"
	"github.com/quocbang/grpc-gateway/pkg/grpc/interceptors"
	"github.com/quocbang/grpc-gateway/server"
	"github.com/quocbang/grpc-gateway/server/repositories/connection"
	senderConn "github.com/quocbang/grpc-gateway/server/sender/connection"
	"github.com/quocbang/grpc-gateway/server/worker"
	"github.com/quocbang/grpc-gateway/server/worker/workersetup"
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

	// init worker pool
	worker, asynqServer := workersetup.RegisterWorker(workersetup.RedisConfig{
		RedisConfig: g.Database.Redis,
	}, sender)
	go startWorkerServer(asynqServer, worker.Processor())

	// set up server info
	ss := server.ServerInfo{
		Repo:                 repo,
		SecretKey:            g.SecretKey,
		AccessTokenLifeTime:  g.AccessTokenLifeTime,
		RefreshTokenLifeTime: g.RefreshTokenLifeTime,
		Worker:               worker,
	}

	// register server info
	server := ss.RegisterServer()
	server.NewServer(s)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", g.GrpcHost, g.GrpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen in grpc, error: %v", err)
	}

	log.Printf("serving grpc host: %s, port %d \n", g.GrpcHost, g.GrpcPort)
	return s.Serve(lis)
}

func startWorkerServer(s *asynq.Server, taskProcessor worker.TaskProcessor) {
	mux := asynq.NewServeMux()
	server.NewWorkerMuxServer(mux, taskProcessor)
	if err := s.Start(mux); err != nil {
		log.Fatalf("failed to start worker server, error: %v", err)
	}
}
