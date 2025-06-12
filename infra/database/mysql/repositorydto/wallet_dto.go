package repositorydto

type InputWalletDto struct {
	Wallet   uint64 `json:"wallet"`
	Exchange uint64 `json:"exchange"`
	Coin     uint64 `json:"coin"`
	Amount   string `json:"amount"`
}
type OutputWalletDto struct {
	Wallet   uint64 `json:"wallet"`
	Exchange uint64 `json:"exchange"`
	Coin     uint64 `json:"coin"`
	Amount   string `json:"amount"`
}

type OutputWalletWithCoinDto struct {
	Wallet   uint64 `json:"wallet"`
	Exchange uint64 `json:"exchange"`
	Coin     uint64 `json:"coin"`
	Amount   string `json:"amount"`
	Symbol   string `json:"symbol"`
	Active   bool   `json:"active"`
}
