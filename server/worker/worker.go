package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/quocbang/grpc-gateway/server/worker/distributor"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
	QueueLow      = "low"
)

type Worker interface {
	Distributor() TaskDistributor
	Processor() TaskProcessor
}

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(context.Context, *distributor.VerifyEmailPayload) (*asynq.Task, error)
}

type TaskProcessor interface {
	ProcessTaskSendVerifyEmail(context.Context, *asynq.Task) error
}
