package controller

import "github.com/mdiaas/processor/internal/app/config"

type ServerController interface {
	Start() error
}

func New(cfg *config.API) ServerController {
	return newServer(cfg)
}
