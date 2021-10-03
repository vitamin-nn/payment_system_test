package model

type User struct {
	ID           string       `json:"id"`
	Email        string    `json:"email"`
	Name    string    `json:"name"`
	City         string    `json:"city"`
	Country string `json:"country"`
	WalletID string `json:"wallet_id"`
}
