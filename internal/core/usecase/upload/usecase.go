package upload

import (
	"context"
	"fmt"

	"github.com/mdiaas/processor/internal/core/dataprovider"
	"github.com/mdiaas/processor/internal/core/entity"
	"github.com/mdiaas/processor/internal/core/entity/errorslabel"
	"github.com/mdiaas/processor/internal/platform/utils"
	log "github.com/sirupsen/logrus"
)

const (
	validFileHeader = "STUDENT_FULL_NAME;STUDENT_DOCUMENT_NUMBER;STUDENT_BIRTH_DATE;STUDENT_MOTHER_NAME;STUDENT_FATHER_NAME;\n"
)

type useCase struct {
	cloudStorage dataprovider.CloudStorage
	publisher    dataprovider.UploadProducer
}

func (u *useCase) Do(ctx context.Context, file entity.File) error {
	l := log.WithContext(ctx).WithField("file", file)

	l.Info("receiving file to upload")
	if err := u.validateFields(file); err != nil {
		l.WithError(err).Error("failed to validate fields")
		return err
	}
	if err := u.cloudStorage.Upload(ctx, file); err != nil {
		l.WithError(err).Error("failed to upload files")
		return fmt.Errorf(errorslabel.UploadFailed)
	}
	message := entity.UploadMessage{
		FileName: file.Name,
	}
	if err := u.publisher.Publish(ctx, message); err != nil {
		l.WithError(err).Error("failed to publish message")
		return fmt.Errorf(errorslabel.PublishFailed)
	}
	return nil
}

func (u *useCase) validateFields(file entity.File) error {
	if utils.IsEmpty(file.Name) {
		return fmt.Errorf(fmt.Sprintf(errorslabel.RequiredField, "file_name"))
	}
	if file.Header != validFileHeader {
		return fmt.Errorf(errorslabel.FileInvalidHeader)
	}
	return nil
}
