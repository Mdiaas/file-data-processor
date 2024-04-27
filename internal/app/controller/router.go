package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mdiaas/processor/internal/app/controller/handlers"
	log "github.com/sirupsen/logrus"
)

const (
	groupPrefix    = "v1/file-data-processor"
	uploadEndpoint = "/upload"
)

type route struct {
	method   string
	endpoint string
	f        func(c echo.Context) error
}

type subRouter struct {
	group  *echo.Group
	routes []*route
}

type router struct {
	fileDataProcessor *subRouter
}

func newRouter(e *echo.Echo) *router {
	handler := handlers.New()
	return &router{
		fileDataProcessor: &subRouter{
			group: e.Group(groupPrefix),
			routes: []*route{
				{
					method:   http.MethodPost,
					endpoint: uploadEndpoint,
					f:        handler.UploadHandler.Upload,
				},
			},
		},
	}
}
func (ref *router) build() {
	for _, route := range ref.fileDataProcessor.routes {
		ref.setRoute(ref.fileDataProcessor.group, route)
	}
}

func (ref *router) setRoute(g *echo.Group, r *route) {
	switch r.method {
	case http.MethodGet:
		g.GET(r.endpoint, r.f)
	case http.MethodPost:
		g.POST(r.endpoint, r.f)
	case http.MethodPut:
		g.PUT(r.endpoint, r.f)
	case http.MethodDelete:
		g.DELETE(r.endpoint, r.f)
	case http.MethodPatch:
		g.PATCH(r.endpoint, r.f)
	default:
		log.Error("method unimplemented")
	}
}
