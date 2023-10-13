package suites

import (
	"fmt"
	"log"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/utils/random"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/repositories/orm/models"
	"github.com/quocbang/grpc-gateway/server/repositories/tests/internal/setup/dependency"
)

type BasicSuite struct {
	*suite.Suite
	Repo       repositories.Repositories
	Models     []models.Model
	DB         *gorm.DB
	ClearTable func() error
}

func NewSuite(bs BasicSuite) *BasicSuite {
	field := []zap.Field{
		zap.String("random seed", random.RandomString(30)),
	}
	logger.Info("start service test", field...)
	return &BasicSuite{
		Models: bs.Models,
		Suite:  &suite.Suite{},
	}
}

func (b *BasicSuite) SetupSuite() {
	repo, err := dependency.InitRepositoriesTest()
	if err != nil {
		logger.Error(fmt.Sprintf("failed to init repositories, error: %v", err))
	}

	database, err := dependency.InitDatabaseConnection()
	if err != nil {
		logger.Error(fmt.Sprintf("failed to init database conn, error: %v", err))
	}

	b.DB = database.PG
	b.Repo = repo
	b.ClearTable = b.Clear
}

func (b *BasicSuite) TearDownSuite() {

}

func (b *BasicSuite) Clear() error {
	for _, m := range b.Models {
		if err := b.DB.Where("1=1").Delete(m).Error; err != nil {
			return err
		}
	}
	return nil
}

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
}
