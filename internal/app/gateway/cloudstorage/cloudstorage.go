package cloudstorage

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/mdiaas/processor/internal/core/dataprovider"
	"github.com/mdiaas/processor/internal/core/entity"
	log "github.com/sirupsen/logrus"
)

type cloudStorage struct {
	client     *storage.Client
	bucketName string
}

func (c *cloudStorage) Upload(ctx context.Context, file entity.File) error {
	l := log.WithContext(ctx).WithField("file", file).WithField("bucket", c.bucketName)
	l.Info("receiving new upload call")
	writer := c.client.Bucket(c.bucketName).Object(file.Name).NewWriter(ctx)
	_, err := io.Copy(writer, file.File)
	if err != nil {
		l.WithError(err).Error(err)
		return err
	}
	if err := writer.Close(); err != nil {
		l.WithError(err).Error(err)
		return err
	}
	return nil
}

func (c *cloudStorage) OpenReader(ctx context.Context, fileName string) (io.Reader, error) {
	return c.client.Bucket(c.bucketName).Object(fileName).NewReader(ctx)
}

func NewCloudStorage(client *storage.Client, bucketName string) dataprovider.CloudStorage {
	return &cloudStorage{
		client:     client,
		bucketName: bucketName,
	}
}
