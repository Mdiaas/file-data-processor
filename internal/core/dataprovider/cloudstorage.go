package dataprovider

import (
	"context"

	"github.com/mdiaas/processor/internal/core/entity"
)

type CloudStorage interface {
	Upload(ctx context.Context, file entity.File) error
}
