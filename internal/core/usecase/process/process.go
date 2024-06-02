package process

import "context"

type Process interface {
	Do(ctx context.Context, fileName string) error
}

func New() Process {
	return &useCase{}
}
