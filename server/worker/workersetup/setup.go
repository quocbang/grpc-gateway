package workersetup

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/quocbang/grpc-gateway/config"
	"github.com/quocbang/grpc-gateway/pkg/grpc/interceptors"
	"github.com/quocbang/grpc-gateway/server/sender"
	"github.com/quocbang/grpc-gateway/server/worker"
	"github.com/quocbang/grpc-gateway/server/worker/logging"
	"go.uber.org/zap"
)

type distributeConfig struct {
	client *asynq.Client
}

type processConfig struct {
	server *asynq.Server
	sender sender.Sender
}

type workerPool struct {
	distribute distributeConfig
	process    processConfig
}

type RedisConfig struct {
	config.RedisConfig
}

func RegisterWorker(rc RedisConfig, sender sender.Sender) (worker.Worker, *asynq.Server) {
	clientOpts := asynq.RedisClientOpt{
		Addr:     rc.Address,
		Password: rc.Password,
	}
	client := asynq.NewClient(clientOpts)

	server := asynq.NewServer(clientOpts, asynq.Config{
		Concurrency: 8, // max my pc concurrency support
		Queues: map[string]int{
			worker.QueueCritical: 10,
			worker.QueueDefault:  5,
			worker.QueueLow:      2,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			maxRetry, _ := asynq.GetMaxRetry(ctx)
			queueName, _ := asynq.GetQueueName(ctx)
			fields := []zap.Field{
				zap.String("typeName", task.Type()),
				zap.Int("maxRetry", maxRetry),
				zap.String("queue", queueName),
				zap.String("payload", string(task.Payload())),
			}
			interceptors.GetLoggerFormContext(ctx).Error("tracing WORKER", fields...)
		}),
		Logger: logging.NewAsynqLog(),
	})

	return workerPool{
		distribute: distributeConfig{
			client: client,
		},
		process: processConfig{
			server: server,
			sender: sender,
		},
	}, server
}
