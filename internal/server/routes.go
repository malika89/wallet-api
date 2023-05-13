package server

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	serviceHandler "wallet-api/internal/handler"
	"wallet-api/internal/svc"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes([]rest.Route{
		{
			Method:  http.MethodPost,
			Path:    "/transfer",
			Handler: serviceHandler.TransferHandler(serverCtx),
		},
	})

}
