package cloudstorage

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/mdiaas/processor/internal/core/dataprovider"
	"github.com/mdiaas/processor/internal/core/entity"
	"github.com/sirupsen/logrus"
)

type cloudStorage struct {
	client     *storage.Client
	bucketName string
}

func (c *cloudStorage) Upload(ctx context.Context, file entity.File) error {
	writer := c.client.Bucket(c.bucketName).Object(file.Name).NewWriter(ctx)
	_, err := io.Copy(writer, file.File)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error(err)
	}
	return nil
}

func NewCloudStorage(client *storage.Client) dataprovider.CloudStorage {
	return &cloudStorage{
		client: client,
	}
}
