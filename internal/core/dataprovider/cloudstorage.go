package dataprovider

import (
	"context"
	"io"

	"github.com/mdiaas/processor/internal/core/entity"
)

type CloudStorage interface {
	Upload(ctx context.Context, file entity.File) error
	OpenReader(ctx context.Context, fileName string) (io.Reader, error)
}
