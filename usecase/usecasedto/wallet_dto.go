package usecasedto

type InputWalletWithCoinDto struct {
	Exchange uint64
	Coin     uint64
}

type OutputWalletWithCoinDto struct {
	Wallet   uint64
	Exchange uint64
	Coin     uint64
	Amount   string
	Symbol   string
	Active   bool
}

type InputCreateWalletDto struct {
	Exchange uint64
	Coin     uint64
	Amount   string
}

type OutputCreateWalletDto struct {
	Wallet   uint64
	Exchange uint64
	Coin     uint64
	Amount   string
}

type InputUpdateWalletDto struct {
	Wallet   uint64
	Exchange uint64
	Coin     uint64
	Amount   string
}

type OutputUpdateWalletDto struct {
	Wallet   uint64
	Exchange uint64
	Coin     uint64
	Amount   string
}
