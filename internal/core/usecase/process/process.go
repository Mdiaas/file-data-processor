package process

import (
	"context"

	"github.com/mdiaas/processor/internal/core/dataprovider"
	"github.com/mdiaas/processor/internal/core/usecase/workerchannel"
)

type Process interface {
	Do(ctx context.Context, fileName string) error
}

func New(storage dataprovider.CloudStorage, worker workerchannel.WorkerChannel, firestorage dataprovider.Firestorage) Process {
	return &useCase{
		storage:       storage,
		workerChannel: worker,
		firestorage:   firestorage,
	}
}
