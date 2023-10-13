package account

import (
	"testing"

	"github.com/quocbang/grpc-gateway/server/tests/setup/suites"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suites.BasicSuite
}

func newSuiteTest() *Suite {
	s := suites.NewSuite()
	return &Suite{BasicSuite: *s}
}

func TestSuite(t *testing.T) {
	suite.Run(t, newSuiteTest())
}
