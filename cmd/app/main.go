package main

import (
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mdiaas/processor/internal/app/config"
	log "github.com/sirupsen/logrus"
)

var cfg config.Config

func init() {
	cfg = config.Load()
}
func main() {
	e := echo.New()
	e.POST("/upload", upload)
	e.Logger.Fatal(e.Start(cfg.API.ServerHost))
}

func upload(c echo.Context) error {
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

	return c.String(http.StatusOK, "Hello, World!")
}
