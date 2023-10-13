package dependency

import (
	"github.com/quocbang/grpc-gateway/server/repositories"
	mocks "github.com/quocbang/grpc-gateway/server/repositories/mock"
	"github.com/stretchr/testify/mock"
)

func NewMockRepositories(m *mock.Mock) repositories.Repositories {
	return &mocks.Repositories{
		Mock: mock.Mock{
			ExpectedCalls: m.ExpectedCalls,
			Calls:         m.Calls,
		},
	}
}
