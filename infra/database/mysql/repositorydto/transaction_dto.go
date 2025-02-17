package repositorydto

type InputTransactionDto struct {
	Transaction     uint64
	TransactionType uint64
	WalletSent      float64
	WalletReceived  float64
	Amount          float64
}

type OutputTransactionDto struct {
	Transaction     uint64
	TransactionType uint64
	WalletSent      float64
	WalletReceived  float64
	Amount          float64
}
