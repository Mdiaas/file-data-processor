package dataprovider

import (
	"context"

	"github.com/mdiaas/processor/internal/core/entity"
)

type Firestorage interface {
	AddStudent(ctx context.Context, student entity.Student) error
}
