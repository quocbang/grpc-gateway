package dependency

import (
	"time"

	"github.com/quocbang/grpc-gateway/server"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/sender"
)

type serverTestOption struct {
	repositories         repositories.Repositories
	sender               sender.Sender
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
	secretKey            string
}

type ServerTestOptions func(*serverTestOption)

func NewMockServer(opts ...ServerTestOptions) server.Server {
	sto := &serverTestOption{}

	for _, opt := range opts {
		opt(sto)
	}

	serverInfo := server.ServerInfo{
		Repo:                 sto.repositories,
		Sender:               sto.sender,
		AccessTokenLifeTime:  sto.accessTokenDuration,
		RefreshTokenLifeTime: sto.refreshTokenDuration,
		SecretKey:            sto.secretKey,
	}
	return serverInfo.RegisterServer()
}

func WithMockRepositories(r repositories.Repositories) ServerTestOptions {
	return func(sto *serverTestOption) {
		sto.repositories = r
	}
}

func WithMockSender(s sender.Sender) ServerTestOptions {
	return func(sto *serverTestOption) {
		sto.sender = s
	}
}

func WithMockAccessTokenDuration(d time.Duration) ServerTestOptions {
	return func(sto *serverTestOption) {
		sto.accessTokenDuration = d
	}
}

func WithMockRefreshTokenDuration(d time.Duration) ServerTestOptions {
	return func(sto *serverTestOption) {
		sto.refreshTokenDuration = d
	}
}

func WithMockSecretKey(secretKey string) ServerTestOptions {
	return func(sto *serverTestOption) {
		sto.secretKey = secretKey
	}
}
