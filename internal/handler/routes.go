package handler

import (
	"go_one/helper/server/http/handler"
	"go_one/internal/registry"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

const (
	BasePrefix = "/one-svc"
	RestPrefix = BasePrefix + "/api/v1"
)

type RestHandler struct {
	svc *registry.ServiceContext
}

func NewRestHandler(svc *registry.ServiceContext) RestHandler {
	return RestHandler{svc: svc}
}
func (h RestHandler) Register(svr *rest.Server) {
	handler.RegisterSwaggerHandler(svr, BasePrefix)
	globalMiddleware(svr, h.svc)
	registerClientHandler(svr, h.svc)

}
func registerClientHandler(svr *rest.Server, svc *registry.ServiceContext) {
	h := NewClientHandler(svc)

	svr.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{
				// svc.AuthMiddleware,
			},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/detail",
					Handler: h.Detail(),
				},
				{
					Method:  http.MethodGet,
					Path:    "/health",
					Handler: h.Health(),
				},
			}...,
		),
		rest.WithPrefix(RestPrefix),
	)

}

func globalMiddleware(_ *rest.Server, _ *registry.ServiceContext) {
}
