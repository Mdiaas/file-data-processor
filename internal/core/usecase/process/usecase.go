package process

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/mdiaas/processor/internal/core/dataprovider"
	"github.com/mdiaas/processor/internal/core/entity"
	"github.com/mdiaas/processor/internal/core/entity/errorslabel"
	"github.com/mdiaas/processor/internal/core/entity/types"
	"github.com/mdiaas/processor/internal/core/usecase/workerchannel"
	"github.com/mdiaas/processor/internal/platform/utils"
	"github.com/sirupsen/logrus"
)

type useCase struct {
	storage       dataprovider.CloudStorage
	workerChannel workerchannel.WorkerChannel
}

func (u *useCase) Do(ctx context.Context, fileName string) error {
	if utils.IsEmpty(fileName) {
		return fmt.Errorf(fmt.Sprintf(errorslabel.RequiredField, "file_name"))
	}
	fileReader, err := u.storage.OpenReader(ctx, fileName)
	if err != nil {
		return fmt.Errorf(errorslabel.ReadFileFailed)
	}
	csv := csv.NewReader(fileReader)
	csv.Comma = ';'
	isHeaderRow := true
	for {
		line, err := csv.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if isHeaderRow {
			isHeaderRow = false
			continue
		}
		u.workerChannel.Add(u.readLine(ctx, fileName, line))
	}
	return nil
}

func (u *useCase) readLine(ctx context.Context, fileName string, line []string) func() {
	l := logrus.WithContext(ctx).WithFields(logrus.Fields{"file_name": fileName, "line": line})
	f := func() {
		l.Info("starting parse line to student")
		student := u.parseLineToStudent(line)
		fmt.Println(student)
	}
	return f
}

func (u *useCase) parseLineToStudent(line []string) entity.Student {
	var errors []string
	birthDate, err := time.Parse(entity.BirthDateFormat, line[types.BirthDateColumn])
	if err != nil {
		birthDate = entity.DefaultDateWhenItsInvalid
		errors = append(errors, fmt.Sprintf("invalid birth date :%s", err.Error()))
	}

	fullName := line[types.FullNameColumn]
	if utils.IsEmpty(fullName) {
		errors = append(errors, "invalid name")
	}

	documentNumber := line[types.DocumentNumberColumn]
	if utils.IsEmpty(documentNumber) {
		errors = append(errors, "invalid documentNumber")
	}

	motherName := line[types.MotherNameColumn]
	if utils.IsEmpty(motherName) {
		errors = append(errors, "invalid name")
	}

	return entity.Student{
		Id:             uuid.NewString(),
		FullName:       fullName,
		DocumentNumber: documentNumber,
		BirthDate:      birthDate,
		MotherName:     motherName,
		FatherName:     line[types.FatherNameColumn],
		Errors:         errors,
	}
}
