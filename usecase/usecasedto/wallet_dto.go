package usecasedto

type InputWalletWithCoinDto struct {
	Exchange uint64
	Coin     uint64
}

type OutputWalletWithCoinDto struct {
	Wallet   uint64
	Exchange uint64
	Coin     uint64
	Amount   float64
	Symbol   string
	Active   bool
}

type InputCreateWalletDto struct {
	Exchange uint64
	Coin     uint64
	Amount   float64
}

type OutputCreateWalletDto struct {
	Wallet   uint64
	Exchange uint64
	Coin     uint64
	Amount   float64
}

type InputUpdateWalletDto struct {
	Wallet   uint64
	Exchange uint64
	Coin     uint64
	Amount   float64
}

type OutputUpdateWalletDto struct {
	Wallet   uint64
	Exchange uint64
	Coin     uint64
	Amount   float64
}
