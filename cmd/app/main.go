package main

import (
	"github.com/mdiaas/processor/internal/app/config"
	"github.com/mdiaas/processor/internal/app/controller"
	log "github.com/sirupsen/logrus"
)

var cfg config.Config

func init() {
	cfg = config.Load()
}
func main() {
	controller := controller.New(&cfg.API)
	if err := controller.Start(); err != nil {
		log.WithError(err).Error("failed to start application")
	}
}

// func upload(c echo.Context) error {
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		log.WithError(err).Error("failed to upload file")
// 		return c.JSON(http.StatusBadRequest, "file is required")
// 	}
// 	src, err := file.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer func(src multipart.File) {
// 		err := src.Close()
// 		if err != nil {
// 			log.WithError(err).Error("failed to close file")
// 		}
// 	}(src)

// 	return c.String(http.StatusOK, "Hello, World!")
// }
