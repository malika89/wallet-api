package svc

import (
	"wallet-api/client"
	"wallet-api/internal/config"
)

type ServiceContext struct {
	c            config.Config
	Transfer     func(req *client.TransferRequest) (*client.TransferResponse, error)
	CreateWallet func(req *client.CreateWalletRequest) (*client.CreateWalletResponse, error)
	GetBalance   func(req *client.GetBalanceRequest) (*client.GetBalanceResponse, error)
	WithDraw     func(req *client.WithdrawReq) (*client.WithdrawResp, error)
	Deposit      func(req *client.DepositReq) (*client.DepositResp, error)
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := &ServiceContext{
		c: c,
	}
	return ctx
}
