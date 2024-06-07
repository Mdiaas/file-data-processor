package process

import (
	"context"

	"github.com/mdiaas/processor/internal/core/dataprovider"
	"github.com/mdiaas/processor/internal/core/usecase/workerchannel"
)

type Process interface {
	Do(ctx context.Context, fileName string) error
}

func New(storage dataprovider.CloudStorage, worker workerchannel.WorkerChannel) Process {
	return &useCase{
		storage:       storage,
		workerChannel: worker,
	}
}
