package handlers

import "github.com/mdiaas/processor/internal/core/usecase/upload"

type Handler struct {
	UploadHandler uploadHandler
}

func New() Handler {
	uploadUC := upload.New()
	uploadHandler := newUploadHandler(uploadUC)
	return Handler{
		UploadHandler: uploadHandler,
	}
}
