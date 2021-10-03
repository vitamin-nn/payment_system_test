package model

type WalletOperation struct {
	ID           string       `json:"id"`
	WalletIDIn        string    `json:"wallet_id_in"`
	WalletIDOut        string    `json:"wallet_id_out"`
	Amount int64 `json:"amount"`
	AmountUSD int64 `json:"amount_usd"`
}
