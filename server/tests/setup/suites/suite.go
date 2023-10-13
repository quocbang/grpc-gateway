package suites

import (
	"log"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/quocbang/data-flow-sync/server/utils/random"
	"github.com/quocbang/grpc-gateway/server"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/sender"
	"github.com/quocbang/grpc-gateway/server/tests/setup/dependency"
)

type ServerConfig struct {
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	SecretKey            string
}

type BasicSuite struct {
	*suite.Suite
	NewMockRepositories func(*mock.Mock) repositories.Repositories
	NewMockSender       func(*mock.Mock) sender.Sender
	NewMockServer       func(...dependency.ServerTestOptions) server.Server
}

func NewSuite() *BasicSuite {
	field := []zap.Field{
		zap.String("random seed", random.RandomString(30)),
	}
	logger.Info("start service test", field...)
	return &BasicSuite{
		NewMockRepositories: dependency.NewMockRepositories,
		NewMockSender:       dependency.NewMockSender,
		NewMockServer:       dependency.NewMockServer,
		Suite:               &suite.Suite{},
	}
}

func (b *BasicSuite) SetupSuite() {
}

func (b *BasicSuite) TearDownSuite() {

}

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
}
