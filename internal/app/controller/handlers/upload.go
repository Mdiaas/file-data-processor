package handlers

import (
	"bufio"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mdiaas/processor/internal/core/entity"
	"github.com/mdiaas/processor/internal/core/usecase/upload"
	log "github.com/sirupsen/logrus"
)

type uploadHandler struct {
	uploadUseCase upload.Upload
}

func (h *uploadHandler) Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		log.WithError(err).Error("failed to upload file")
		return c.JSON(http.StatusBadRequest, "file is required")
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			log.WithError(err).Error("failed to close file")
		}
	}(src)
	fileHeader, err := extractHeaderFromFile(file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error to read file header")
	}
	fileEntity := entity.File{
		Header: fileHeader,
		Name:   file.Filename,
		File:   src,
	}
	if err := h.uploadUseCase.Do(c.Request().Context(), fileEntity); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func extractHeaderFromFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return line, nil

}
func newUploadHandler(uploadUseCase upload.Upload) uploadHandler {
	return uploadHandler{
		uploadUseCase: uploadUseCase,
	}
}
