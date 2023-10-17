package email

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"

	"github.com/quocbang/grpc-gateway/server/worker"
	"github.com/quocbang/grpc-gateway/server/worker/distributor"
)

const (
	TaskSendVerifyEmail = "Task:SendVerifyEmail"
)

type emailDistributor struct {
	client *asynq.Client
}

func NewEmailDistributor(client *asynq.Client) worker.TaskDistributor {
	return &emailDistributor{
		client: client,
	}
}

func (rd *emailDistributor) DistributeTaskSendVerifyEmail(ctx context.Context, payload *distributor.VerifyEmailPayload) (*asynq.Task, error) {
	p, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload, error: %v", err)
	}

	// create new task and enqueue task to redis.
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(1 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	task := asynq.NewTask(TaskSendVerifyEmail, p, opts...)
	_, err = rd.client.Enqueue(task)
	if err != nil {
		return nil, fmt.Errorf("failed to enqueue send verify email task, error: %v", err)
	}

	return task, nil
}
