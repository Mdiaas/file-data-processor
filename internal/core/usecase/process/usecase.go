package process

import (
	"context"
	"fmt"
)

type useCase struct {
}

func (u *useCase) Do(ctx context.Context, fileName string) error {
	fmt.Println("here ", fileName)
	return nil
}
