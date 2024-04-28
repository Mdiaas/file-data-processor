package main

import (
	"github.com/mdiaas/processor/internal/app/config"
	"github.com/mdiaas/processor/internal/app/controller"
	log "github.com/sirupsen/logrus"
)

var cfg config.Config

func init() {
	cfg = config.Load()
}
func main() {
	controller := controller.New(&cfg)
	if err := controller.Start(); err != nil {
		log.WithError(err).Error("failed to start application")
	}
}
