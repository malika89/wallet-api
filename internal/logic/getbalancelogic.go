package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"wallet-api/client"
	"wallet-api/internal/svc"
)

type GetBalanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBalanceLogic {
	return &GetBalanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBalanceLogic) GetBalance(in *client.GetBalanceRequest) (*client.GetBalanceResponse, error) {
	// 实现查询余额的逻辑，这里只是返回一个示例余额
	return &client.GetBalanceResponse{
		Balance: 1000.0,
	}, nil
}
