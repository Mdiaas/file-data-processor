package upload

import (
	"context"

	"github.com/mdiaas/processor/internal/core/dataprovider"
	"github.com/mdiaas/processor/internal/core/entity"
)

type Upload interface {
	Do(ctx context.Context, file entity.File) error
}

func New(cloudStorage dataprovider.CloudStorage, publisher dataprovider.UploadProducer) Upload {
	return &useCase{
		cloudStorage: cloudStorage,
		publisher:    publisher,
	}
}
