package main

import (
	"context"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"github.com/mdiaas/processor/internal/app/config"
	"github.com/mdiaas/processor/internal/app/controller"
	"github.com/mdiaas/processor/internal/app/controller/listener"
	"github.com/mdiaas/processor/internal/app/gateway/cloudstorage"
	"github.com/mdiaas/processor/internal/core/usecase/process"
	"github.com/mdiaas/processor/internal/core/usecase/workerchannel"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

var cfg config.Config

func init() {
	cfg = config.Load()
}
func main() {
	ctx := context.Background()
	opt := option.WithCredentialsFile(cfg.Google.Credentials)
	client, err := pubsub.NewClient(ctx, cfg.Project.Id, opt)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	storageClient, err := storage.NewClient(ctx, opt)
	if err != nil {
		log.WithContext(ctx).WithError(err).Fatal("failed to create pubsub client")
	}
	cloudstorage := cloudstorage.NewCloudStorage(storageClient, cfg.CloudStorage.BucketName)
	workerUc := workerchannel.New(cfg.Workers.NumberOfWorks)
	processUc := process.New(cloudstorage, workerUc)
	listener := listener.New(client, &cfg, processUc)
	go listener.Listen(ctx)
	controller := controller.New(&cfg)
	if err := controller.Start(); err != nil {
		log.WithError(err).Error("failed to start application")
	}
}
