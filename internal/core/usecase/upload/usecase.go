package upload

import (
	"context"

	"github.com/mdiaas/processor/internal/core/entity"
)

type useCase struct{}

func (u *useCase) Do(ctx context.Context, file entity.File) error {
	return nil
}
