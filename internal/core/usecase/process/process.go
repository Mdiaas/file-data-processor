package process

import (
	"context"

	"github.com/mdiaas/processor/internal/core/dataprovider"
)

type Process interface {
	Do(ctx context.Context, fileName string) error
}

func New(storage dataprovider.CloudStorage) Process {
	return &useCase{
		storage: storage,
	}
}
