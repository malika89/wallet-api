package handler

import (
	"context"
	"github.com/malika89/wallet-api/client"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/malika89/wallet-api/internal/logic"
	"github.com/malika89/wallet-api/internal/svc"
)

func TransferHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req client.TransferRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		ctx_, cancel := context.WithCancel(context.Background())
		defer cancel()
		l := logic.NewTransferLogic(ctx_, ctx)
		resp, err := l.Transfer(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
