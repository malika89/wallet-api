package internal

type CreateWalletRequest struct {
	Name string
}

type CreateWalletResponse struct {
	Address string
}

type GetBalanceRequest struct {
	Address string
}

type GetBalanceResponse struct {
	Balance string
}

type WithdrawRequest struct {
	Address string
	Amount  string
}

type WithdrawResponse struct {
	TxHash string
}

type DepositRequest struct {
	Address string
	Amount  string
}

type DepositResponse struct {
	TxHash string
}
