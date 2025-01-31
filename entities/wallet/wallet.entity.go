package wallet

type Wallet struct {
	Wallet   int64   `json:"wallet"`
	Exchange int64   `json:"exchange"`
	User     int64   `json:"user"`
	Coin     int64   `json:"coin"`
	Amount   float64 `json:"amount"`
}

type WalletWithCoin struct {
	Wallet   int64   `json:"wallet"`
	Exchange int64   `json:"exchange"`
	User     int64   `json:"user"`
	Amount   float64 `json:"amount"`
	Coin     int64   `json:"coin"`
	Symbol   string  `json:"symbol"`
	Active   bool    `json:"active"`
}
