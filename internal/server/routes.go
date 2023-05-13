package server

import (
	serviceHandler "github.com/malika89/wallet-api/internal/handler"
	"github.com/malika89/wallet-api/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
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
