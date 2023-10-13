package suites

import (
	"fmt"
	"log"

	"github.com/quocbang/data-flow-sync/server/utils/random"
	"github.com/quocbang/grpc-gateway/server/sender"
	"github.com/quocbang/grpc-gateway/server/sender/tests/setup/dependency"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type BasicSuite struct {
	*suite.Suite
	Sender sender.Sender
}

func NewSuite() *BasicSuite {
	field := []zap.Field{
		zap.String("random seed", random.RandomString(30)),
	}
	logger.Info("start service test", field...)
	return &BasicSuite{
		Suite: &suite.Suite{},
	}
}

func (b *BasicSuite) SetupSuite() {
	sender, err := dependency.InitSenderTest()
	if err != nil {
		logger.Error(fmt.Sprintf("failed to init sender, error: %v", err))
	}
	b.Sender = sender
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
