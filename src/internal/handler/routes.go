// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2

package handler

import (
	"net/http"

	"template/gozero/src/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// Phan Chinh Quan
				Method:  http.MethodGet,
				Path:    "/health",
				Handler: HealthHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
		},
	)
}
