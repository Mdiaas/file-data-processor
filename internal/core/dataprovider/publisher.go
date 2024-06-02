package dataprovider

import (
	"context"

	"github.com/mdiaas/processor/internal/core/entity"
)

type UploadProducer interface {
	Publish(ctx context.Context, message entity.UploadMessage) error
}
