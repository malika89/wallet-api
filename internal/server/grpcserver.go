package server

import (
	"context"
	"wallet-api/client"
	"wallet-api/internal/logic"
	"wallet-api/internal/svc"
	"wallet-api/proto"
)

type WalletServer struct {
	proto.UnimplementedWalletServer
	svcCtx *svc.ServiceContext
}

func NewWalletServer(svcCtx *svc.ServiceContext) *WalletServer {
	return &WalletServer{
		svcCtx: svcCtx,
	}
}

func (s *WalletServer) CreateWallet(ctx context.Context, req *client.CreateWalletRequest) (*client.CreateWalletResponse, error) {
	res, err := s.svcCtx.CreateWallet(req)
	if err != nil {
		return nil, err
	}
	return &client.CreateWalletResponse{
		Address: res.Address,
	}, nil
}

func (s *WalletServer) GetBalance(ctx context.Context, req *client.GetBalanceRequest) (*client.GetBalanceResponse, error) {
	l := logic.NewGetBalanceLogic(ctx, s.svcCtx)
	return l.GetBalance(req)
}

func (s *WalletServer) Deposit(ctx context.Context, req *client.DepositReq) (*client.DepositResp, error) {
	res, err := s.svcCtx.Deposit(req)
	if err != nil {
		return nil, err
	}
	return &client.DepositResp{
		TxHash: res.TxHash,
	}, nil
}

func (s *WalletServer) WithDraw(ctx context.Context, req *client.WithdrawReq) (*client.WithdrawResp, error) {
	res, err := s.svcCtx.WithDraw(req)
	if err != nil {
		return nil, err
	}
	return &client.WithdrawResp{
		TxHash: res.TxHash,
	}, nil
}
