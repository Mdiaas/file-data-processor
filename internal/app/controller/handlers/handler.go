package handlers

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/mdiaas/processor/internal/app/config"
	"github.com/mdiaas/processor/internal/app/gateway/cloudstorage"
	"github.com/mdiaas/processor/internal/core/usecase/upload"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

type Handler struct {
	UploadHandler uploadHandler
}

func New(cfg *config.Config) Handler {
	ctx := context.Background()
	opt := option.WithCredentialsFile(cfg.Google.Credentials)
	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		log.WithContext(ctx).WithError(err).Fatal("failed to create storage client")
	}
	cloudstorage := cloudstorage.NewCloudStorage(client, cfg.CloudStorage.BucketName)
	uploadUC := upload.New(cloudstorage)
	uploadHandler := newUploadHandler(uploadUC)
	return Handler{
		UploadHandler: uploadHandler,
	}
}
