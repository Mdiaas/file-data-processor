package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mdiaas/processor/internal/core/usecase/upload"
	log "github.com/sirupsen/logrus"
)

type uploadHandler struct {
	uploadUseCase upload.Upload
}

func (h *uploadHandler) Upload(c echo.Context) error {
	log.Info("receiving new upload call")
	return c.NoContent(http.StatusOK)
}

func newUploadHandler(uploadUseCase upload.Upload) uploadHandler {
	return uploadHandler{
		uploadUseCase: uploadUseCase,
	}
}
