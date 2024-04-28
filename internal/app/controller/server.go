package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mdiaas/processor/internal/app/config"
	log "github.com/sirupsen/logrus"
)

type server struct {
	cfg    *config.Config
	e      *echo.Echo
	router *router
}

func newServer(cfg *config.Config) *server {
	srv := new(server)
	srv.cfg = cfg
	srv.e = echo.New()
	router := newRouter(srv.e, cfg)
	srv.router = router
	return srv
}

func (ref *server) Start() error {
	log.WithField("serverHost", ref.cfg.API.ServerHost).Info("starting aplication")
	ref.router.build()
	return ref.e.Start(ref.cfg.API.ServerHost)
}
