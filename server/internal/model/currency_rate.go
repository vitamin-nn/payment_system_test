package model

type CurrencyRate struct {
	CurrencyCode string `json:"currency_code"`
	Rate int64 `json:"rate"` // exchange currency: 100 curr to 100 USD
}
