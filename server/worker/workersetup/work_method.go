package workersetup

import (
	"github.com/quocbang/grpc-gateway/server/worker"
	dEmail "github.com/quocbang/grpc-gateway/server/worker/distributor/email"
	pEmail "github.com/quocbang/grpc-gateway/server/worker/processor/email"
)

func (w workerPool) Distributor() worker.TaskDistributor {
	return dEmail.NewEmailDistributor(w.distribute.client)
}

func (w workerPool) Processor() worker.TaskProcessor {
	return pEmail.NewEmailProcessor(w.process.server, w.process.sender)
}
