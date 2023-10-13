package dependency

import (
	"github.com/quocbang/grpc-gateway/server/sender"
	mocks "github.com/quocbang/grpc-gateway/server/sender/mock"
	"github.com/stretchr/testify/mock"
)

func NewMockSender(m *mock.Mock) sender.Sender {
	return &mocks.Sender{
		Mock: mock.Mock{
			ExpectedCalls: m.ExpectedCalls,
			Calls:         m.Calls,
		},
	}
}
