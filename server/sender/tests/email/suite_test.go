package email

import (
	"testing"

	"github.com/quocbang/grpc-gateway/server/sender/tests/setup/suites"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suites.BasicSuite
}

func NewSuite() *Suite {
	s := suites.NewSuite()
	return &Suite{BasicSuite: *s}
}

func TestSuite(t *testing.T) {
	suite.Run(t, NewSuite())
}
