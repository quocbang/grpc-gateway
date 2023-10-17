package email

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"

	"github.com/quocbang/grpc-gateway/server/sender"
	"github.com/quocbang/grpc-gateway/server/worker"
	"github.com/quocbang/grpc-gateway/server/worker/distributor"
)

type emailProcessor struct {
	server *asynq.Server
	sender sender.Sender
}

func NewEmailProcessor(server *asynq.Server, sender sender.Sender) worker.TaskProcessor {
	return &emailProcessor{
		server: server,
		sender: sender,
	}
}

func (rp *emailProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	verifyEmailPayload := distributor.VerifyEmailPayload{}
	if err := json.Unmarshal(task.Payload(), &verifyEmailPayload); err != nil {
		return fmt.Errorf("failed to unmarshal payload, error: %v %w", err, asynq.SkipRetry)
	}

	err := rp.sender.Email().SendVerifyEmail(ctx, verifyEmailPayload.To, verifyEmailPayload.Subject, verifyEmailPayload.Content)
	if err != nil {
		return fmt.Errorf("failed to send verify email, error: %v", err)
	}

	return nil
}
