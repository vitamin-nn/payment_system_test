package model

type Wallet struct {
	ID           string       `json:"id"`
	CurrencyCode        string    `json:"currency_code"`
	Amount int64 `json:"amount"`
}
