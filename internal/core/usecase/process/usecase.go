package process

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"

	"github.com/mdiaas/processor/internal/core/dataprovider"
	"github.com/mdiaas/processor/internal/core/entity/errorslabel"
	"github.com/mdiaas/processor/internal/platform/utils"
)

type useCase struct {
	storage dataprovider.CloudStorage
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
	for {
		line, err := csv.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		fmt.Println("line: ", line)
	}
	return nil
}
