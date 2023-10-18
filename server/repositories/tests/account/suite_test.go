package email

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/quocbang/grpc-gateway/server/repositories/orm/models"
	"github.com/quocbang/grpc-gateway/server/repositories/tests/internal/setup/suites"
)

type Suite struct {
	suites.BasicSuite
}

func NewSuite() *Suite {
	s := suites.NewSuite(suites.BasicSuite{
		Models: []models.Model{&models.VerifyAccount{}, &models.Account{}}, // the fk should clear first
	})
	return &Suite{
		BasicSuite: *s,
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, NewSuite())
}
